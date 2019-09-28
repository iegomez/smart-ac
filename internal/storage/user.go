package storage

import (
	"database/sql"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// saltSize defines the salt size
const saltSize = 16

// defaultSessionTTL defines the default session TTL. Set default to a year for testing.
const defaultSessionTTL = time.Hour * (24 * 365)

// Any upper, lower, digit characters, at least 6 characters.
var usernameValidator = regexp.MustCompile(`.+@.+`)

// Any printable characters, at least 6 characters
var passwordValidator = regexp.MustCompile(`^.{5,}$`)

// Must contain @ (this is far from perfect)
var emailValidator = regexp.MustCompile(`.+@.+`)

// Validation token TTL (in seconds)
var validationTTL = 60 * 60 * 72

// User represents a user to external code.
type User struct {
	ID           int64     `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password"`
	SessionTTL   int32     `db:"session_ttl"`
	IsActive     bool      `db:"is_active"`
	IsAdmin      bool      `db:"is_admin"`
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
func CreateUser(db *sqlx.DB, user *User, password string) (int64, error) {
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

	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert the new user.
	err = db.Get(&user.ID, `
		insert into "user" (
			created_at,
			updated_at,
			username,
			password_hash,
			session_ttl,
			is_admin
		)
		values (
			$1, $2, $3, $4, $5, $6) returning id`,
		user.CreatedAt,
		user.UpdatedAt,
		user.Username,
		pwHash,
		user.SessionTTL,
		user.IsAdmin,
	)
	if err != nil {
		return 0, handlePSQLError(Insert, err, "insert error")
	}

	log.WithFields(log.Fields{
		"username":    user.Username,
		"session_ttl": user.SessionTTL,
	}).Info("user created")
	return user.ID, nil
}

// GetUser returns the User for the given id.
func GetUser(db sqlx.Queryer, id int64) (User, error) {
	var user User
	err := sqlx.Get(db, &user, "select * from \"user\" where id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrDoesNotExist
		}
		return user, errors.Wrap(err, "select error")
	}

	return user, nil
}

// GetUserCount returns the total number of users.
func GetUserCount(db *sqlx.DB, search string) (int32, error) {
	var count int32
	if search != "" {
		search = search + "%"
	}
	err := db.Get(&count, `
		select
			count(id)
		from "user"
		where
			($1 != '' and username like $1)
			or ($1 = '')
		`, search)
	if err != nil {
		return 0, errors.Wrap(err, "select error")
	}
	return count, nil
}

// GetUsers returns a slice of users, respecting the given limit and offset.
func GetUsers(db *sqlx.DB, limit, offset int32, search string) ([]User, error) {
	var users []User
	if search != "" {
		search = search + "%"
	}
	err := db.Select(&users, `select * from "user" where ($3 != '' and username like $3) or ($3 = '') order by username limit $1 offset $2`, limit, offset, search)
	if err != nil {
		return nil, errors.Wrap(err, "select error")
	}
	return users, nil
}

// UpdateUser updates the given User.
func UpdateUser(db *sqlx.DB, user User) error {
	if err := ValidateUsername(user.Username); err != nil {
		return errors.Wrap(err, "validation error")
	}

	res, err := db.Exec(`
		update "user"
		set
			username = $2,
			session_ttl = $3,
			updated_at = now(),
			is_admin = $4
		where id = $1`,
		user.ID,
		user.Username,
		user.SessionTTL,
		user.IsAdmin,
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
		"id":          user.ID,
		"username":    user.Username,
		"session_ttl": user.SessionTTL,
	}).Info("user updated")

	return nil
}

// DeleteUser deletes the User record matching the given ID.
func DeleteUser(db *sqlx.DB, id int64) error {

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

// LoginUser returns a JWT token for the user matching the given email
// and password.
func LoginUser(db *sqlx.DB, email string, password string) (string, error) {
	// Find the user by email
	var user User
	err := db.Get(&user, "select * from \"user\" where email = $1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrDoesNotExist
		}
		return "", errors.Wrap(err, "select error")
	}

	// Compare the passed in password with the hash in the database.
	if checkPasswordHash(password, user.PasswordHash) {
		log.Errorf("couldn't match %s and %s", password, user.PasswordHash)
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
		"iss": "smart-ac",
		"aud": "smart-ac",
		"nbf": nowSecondsSinceEpoch,
		"exp": expSecondsSinceEpoch,
		"sub": user.Username,
	})

	jwt, err := token.SignedString(jwtsecret)
	if nil != err {
		return jwt, errors.Wrap(err, "get jwt signed string error")
	}

	return jwt, err
}

// UpdatePassword updates the user with the new password.
func UpdatePassword(db *sqlx.DB, id int64, oldpassword, newpassword string) error {
	if err := ValidatePassword(newpassword); err != nil {
		return errors.Wrap(err, "validation error")
	}

	user, err := GetUser(db, id)
	if err != nil {
		return err
	}

	if checkPasswordHash(oldpassword, user.PasswordHash) {
		return ErrInvalidUsernameOrPassword
	}

	pwHash, err := hashPassword(newpassword)
	if err != nil {
		return err
	}

	// Add the new user.
	_, err = db.Exec("update \"user\" set password = $1, updated_at = now() where id = $2",
		pwHash, id)
	if err != nil {
		return errors.Wrap(err, "update error")
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("user password updated")
	return nil

}
