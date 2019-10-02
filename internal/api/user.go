package api

import (
	"fmt"
	"regexp"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/iegomez/smart-ac/api"
	"github.com/pkg/errors"

	"github.com/iegomez/smart-ac/internal/api/helpers"
	"github.com/iegomez/smart-ac/internal/storage"

	log "github.com/sirupsen/logrus"
)

var validAuthorizationRegexp = regexp.MustCompile(`(?i)^bearer (.*)$`)

// UserAPI exports the User related functions.
type UserAPI struct {
}

// NewUserAPI creates a new UserAPI.
func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

func getClaims(ctx context.Context) (*Claims, error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get token from context error")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return storage.JWTSecret(), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt parse error")
	}

	if !token.Valid {
		return nil, errors.Wrap(ErrNotAuthorized, "invalid jwt token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		// no need to use a static error, this should never happen
		return nil, fmt.Errorf("api/auth: expected *Claims, got %T", token.Claims)
	}

	return claims, nil
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", storage.ErrInvalidUsernameOrPassword
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", storage.ErrInvalidUsernameOrPassword
	}

	match := validAuthorizationRegexp.FindStringSubmatch(token[0])

	// authorization header should respect RFC1945
	if len(match) == 0 {
		log.Warning("RFC1945 format expected : Authorization: <type> <credentials>")
		return token[0], nil
	}

	return match[1], nil
}

func GetIsAdmin(ctx context.Context) (bool, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return false, err
	}

	user, err := storage.GetUserByUsername(storage.DB(), claims.Username)
	if err != nil {
		return false, storage.ErrInvalidUsernameOrPassword
	}

	return user.IsAdmin, nil
}

func GetIsUser(ctx context.Context) (bool, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return false, err
	}

	_, err = storage.GetUserByUsername(storage.DB(), claims.Username)
	if err != nil {
		return false, storage.ErrInvalidUsernameOrPassword
	}

	return true, nil
}

func GetMatchesUser(ctx context.Context, id int64) (bool, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return false, err
	}

	user, err := storage.GetUserByUsername(storage.DB(), claims.Username)
	if err != nil {
		return false, storage.ErrInvalidUsernameOrPassword
	}

	if !user.IsAdmin && user.ID != id {
		return false, storage.ErrInvalidUsernameOrPassword
	}

	return true, nil
}

func GetUsername(ctx context.Context) (string, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return "", err
	}

	return claims.Username, nil
}

// Create creates the given user.
func (a *UserAPI) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.User == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "user must not be nil")
	}

	user := storage.User{
		Username:   req.User.Username,
		SessionTTL: req.User.SessionTtl,
		IsAdmin:    req.User.IsAdmin,
	}

	isAdmin, err := GetIsAdmin(ctx)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	if !isAdmin {
		// non-admin users are not able to modify the fields below
		user.IsAdmin = false
		user.SessionTTL = 0
	}

	var userID int64

	err = storage.Transaction(func(tx sqlx.Ext) error {
		userID, err = storage.CreateUser(tx, &user, req.Password)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &pb.CreateUserResponse{Id: userID}, nil
}

// Get returns the user matching the given ID.
func (a *UserAPI) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	user, err := storage.GetUser(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	resp := pb.GetUserResponse{
		User: &pb.User{
			Id:         user.ID,
			Username:   user.Username,
			SessionTtl: user.SessionTTL,
			IsAdmin:    user.IsAdmin,
		},
	}

	resp.CreatedAt, err = ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	resp.UpdatedAt, err = ptypes.TimestampProto(user.UpdatedAt)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &resp, nil
}

// List lists the users.
func (a *UserAPI) List(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	users, err := storage.GetUsers(storage.DB(), int(req.Limit), int(req.Offset), req.Search)
	if err != nil {
		log.Errorf("failed at get users: %s", err)
		return nil, helpers.ErrToRPCError(err)
	}

	totalUserCount, err := storage.GetUserCount(storage.DB(), req.Search)
	if err != nil {
		log.Errorf("failed at get users count: %s", err)
		return nil, helpers.ErrToRPCError(err)
	}

	resp := pb.ListUserResponse{
		TotalCount: int64(totalUserCount),
	}

	for _, u := range users {
		row := pb.UserListItem{
			Id:         u.ID,
			Username:   u.Username,
			SessionTtl: u.SessionTTL,
			IsAdmin:    u.IsAdmin,
		}

		row.CreatedAt, err = ptypes.TimestampProto(u.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		row.UpdatedAt, err = ptypes.TimestampProto(u.UpdatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		resp.Result = append(resp.Result, &row)
	}

	return &resp, nil
}

// Update updates the given user.
func (a *UserAPI) Update(ctx context.Context, req *pb.UpdateUserRequest) (*empty.Empty, error) {
	if req.User == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "user must not be nil")
	}

	isUser, err := GetMatchesUser(ctx, req.User.Id)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	userUpdate := storage.UserUpdate{
		ID:         req.User.Id,
		Username:   req.User.Username,
		IsAdmin:    req.User.IsAdmin,
		SessionTTL: req.User.SessionTtl,
	}

	err = storage.UpdateUser(storage.DB(), userUpdate)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// Delete deletes the user matching the given ID.
func (a *UserAPI) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {

	isAdmin, err := GetIsAdmin(ctx)
	if err != nil || !isAdmin {
		return nil, helpers.ErrToRPCError(err)
	}

	err = storage.DeleteUser(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// UpdatePassword updates the password for the user matching the given ID.
func (a *UserAPI) UpdatePassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*empty.Empty, error) {

	isUser, err := GetMatchesUser(ctx, req.UserId)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	err = storage.UpdatePassword(storage.DB(), req.UserId, req.Password)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// Login validates the login request and returns a JWT token.
func (a *UserAPI) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	jwt, err := storage.LoginUser(storage.DB(), req.Username, req.Password)
	if nil != err {
		log.Errorf("fail login: %s\n", err)
		return nil, helpers.ErrToRPCError(err)
	}

	return &pb.LoginResponse{Jwt: jwt}, nil
}

type claims struct {
	Username string `json:"username"`
}

// Profile returns the user profile.
func (a *UserAPI) Profile(ctx context.Context, req *empty.Empty) (*pb.ProfileResponse, error) {

	username, err := GetUsername(ctx)
	if nil != err {
		log.Errorf("fail 1: %s\n", err)
		return nil, helpers.ErrToRPCError(err)
	}

	// Get the user id based on the username.
	user, err := storage.GetUserByUsername(storage.DB(), username)
	if nil != err {
		log.Errorf("fail 2: %s\n", err)
		return nil, helpers.ErrToRPCError(err)
	}

	resp := pb.ProfileResponse{
		User: &pb.User{
			Id:         user.ID,
			Username:   user.Username,
			SessionTtl: user.SessionTTL,
			IsAdmin:    user.IsAdmin,
		},
	}

	return &resp, nil
}
