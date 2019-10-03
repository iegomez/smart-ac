package storage

import (
	"database/sql"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// saltSize defines the salt size
const saltSize = 16

// defaultSessionTTL defines the default session TTL
const defaultSessionTTL = time.Hour * 24

// Any upper, lower, digit characters, at least 6 characters.
var usernameValidator = regexp.MustCompile(`^[[:alnum:]]+$`)

// Any printable characters, at least 6 characters.
var passwordValidator = regexp.MustCompile(`^.{6,}$`)

// Must contain @ (this is far from perfect)
var emailValidator = regexp.MustCompile(`.+@.+`)

// User represents a user to external code.
type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	IsAdmin      bool      `db:"is_admin"`
	SessionTTL   int32     `db:"session_ttl"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PasswordHash string    `db:"password_hash"`
}

const externalUserFields = "id, username, is_admin, session_ttl, created_at, updated_at"
const internalUserFields = "*"

// UserUpdate represents the user fields that can be "updated" in the simple
// case.  This excludes id, which identifies the record to be updated.
type UserUpdate struct {
	ID         int64  `db:"id"`
	Username   string `db:"username"`
	IsAdmin    bool   `db:"is_admin"`
	SessionTTL int32  `db:"session_ttl"`
}

// userInternal represents a user as known by the database.
type userInternal struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	IsAdmin      bool      `db:"is_admin"`
	SessionTTL   int32     `db:"session_ttl"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// ValidateUsername validates the given username.
func ValidateUsername(username string) error {
	if !usernameValidator.MatchString(username) {
		return ErrUserInvalidUsername
	}
	return nil
}

// ValidatePassword validates the given password.
func ValidatePassword(password string) error {
	if !passwordValidator.MatchString(password) {
		return ErrUserPasswordLength
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateUser creates the given user.
func CreateUser(db sqlx.Queryer, user *User, password string) (int64, error) {
	if err := ValidateUsername(user.Username); err != nil {
		return 0, errors.Wrap(err, "validation error")
	}

	if err := ValidatePassword(password); err != nil {
		return 0, errors.Wrap(err, "validation error")
	}

	pwHash, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	log.Infof("password hash: %s", pwHash)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Add the new user.
	err = sqlx.Get(db, &user.ID, `
		insert into "user" (
			username,
			password_hash,
			is_admin,
			session_ttl,
			created_at,
			updated_at
		)
		values (
			$1, $2, $3, $4, $5, $6) returning id`,
		user.Username,
		pwHash,
		user.IsAdmin,
		user.SessionTTL,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return 0, handlePSQLError(Insert, err, "insert error")
	}

	log.WithFields(log.Fields{
		"username":    user.Username,
		"session_ttl": user.SessionTTL,
		"is_admin":    user.IsAdmin,
	}).Info("user created")
	return user.ID, nil
}

// GetUser returns the User for the given id.
func GetUser(db sqlx.Queryer, id int64) (User, error) {
	var user User
	err := sqlx.Get(db, &user, "select "+externalUserFields+" from \"user\" where id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrDoesNotExist
		}
		return user, errors.Wrap(err, "select error")
	}

	return user, nil
}

// GetUserByUsername returns the User for the given username.
func GetUserByUsername(db sqlx.Queryer, username string) (User, error) {
	var user User
	err := sqlx.Get(db, &user, "select "+externalUserFields+" from \"user\" where username = $1", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrDoesNotExist
		}
		return user, errors.Wrap(err, "select error")
	}

	return user, nil
}

// GetUserCount returns the total number of users.
func GetUserCount(db sqlx.Queryer, search string) (int32, error) {
	var count int32
	if search != "" {
		search = "%" + search + "%"
	}
	err := sqlx.Get(db, &count, `
		select
			count(*)
		from "user"
		where
			($1 != '' and username ilike $1)
			or ($1 = '')
		`, search)
	if err != nil {
		return 0, errors.Wrap(err, "select error")
	}
	return count, nil
}

// GetUsers returns a slice of users, respecting the given limit and offset.
func GetUsers(db sqlx.Queryer, limit, offset int, search string) ([]User, error) {
	var users []User
	if search != "" {
		search = "%" + search + "%"
	}
	err := sqlx.Select(db, &users, "select "+externalUserFields+` from "user" where ($3 != '' and username ilike $3) or ($3 = '') order by username limit $1 offset $2`, limit, offset, search)
	if err != nil {
		return nil, errors.Wrap(err, "select error")
	}
	return users, nil
}

// UpdateUser updates the given User.
func UpdateUser(db sqlx.Execer, item UserUpdate) error {
	if err := ValidateUsername(item.Username); err != nil {
		return errors.Wrap(err, "validation error")
	}

	res, err := db.Exec(`
		update "user"
		set
			username = $2,
			is_admin = $3,
			session_ttl = $4,
			updated_at = now()
		where id = $1`,
		item.ID,
		item.Username,
		item.IsAdmin,
		item.SessionTTL,
	)
	if err != nil {
		return handlePSQLError(Update, err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"id":          item.ID,
		"username":    item.Username,
		"is_admin":    item.IsAdmin,
		"session_ttl": item.SessionTTL,
	}).Info("user updated")

	return nil
}

// DeleteUser deletes the User record matching the given ID.
func DeleteUser(db sqlx.Execer, id int64) error {
	res, err := db.Exec("delete from \"user\" where id = $1", id)
	if err != nil {
		return errors.Wrap(err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("user deleted")
	return nil
}

// LoginUser returns a JWT token for the user matching the given username
// and password.
func LoginUser(db sqlx.Queryer, username string, password string) (string, error) {
	// Find the user by username
	var user userInternal
	err := sqlx.Get(db, &user, "select "+internalUserFields+" from \"user\" where username = $1", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrInvalidUsernameOrPassword
		}
		return "", errors.Wrap(err, "select error")
	}

	// Compare the passed in password with the hash in the database.
	if !checkPasswordHash(password, user.PasswordHash) {
		return "", ErrInvalidUsernameOrPassword
	}

	// Generate the token.
	now := time.Now()
	nowSecondsSinceEpoch := now.Unix()
	var expSecondsSinceEpoch int64
	if user.SessionTTL > 0 {
		expSecondsSinceEpoch = nowSecondsSinceEpoch + (60 * int64(user.SessionTTL))
	} else {
		expSecondsSinceEpoch = nowSecondsSinceEpoch + int64(defaultSessionTTL/time.Second)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "smart-ac",
		"aud":      "smart-ac",
		"nbf":      nowSecondsSinceEpoch,
		"exp":      expSecondsSinceEpoch,
		"sub":      "user",
		"username": user.Username,
	})

	jwt, err := token.SignedString(jwtsecret)
	if nil != err {
		return jwt, errors.Wrap(err, "get jwt signed string error")
	}
	return jwt, err
}

// UpdatePassword updates the user with the new password.
func UpdatePassword(db sqlx.Execer, id int64, newpassword string) error {
	if err := ValidatePassword(newpassword); err != nil {
		return errors.Wrap(err, "validation error")
	}

	pwHash, err := hashPassword(newpassword)
	if err != nil {
		return err
	}

	// Add the new user.
	_, err = db.Exec("update \"user\" set password_hash = $1, updated_at = now() where id = $2",
		pwHash, id)
	if err != nil {
		return errors.Wrap(err, "update error")
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("user password updated")
	return nil

}
