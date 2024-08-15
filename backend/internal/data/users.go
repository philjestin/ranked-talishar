package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *db.Queries
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

var AnonymousUser = &User{}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (m UserModel) Insert(user *User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now()
	args := &db.CreateUserParams{
		UserName:       user.Name,
		UserEmail:      user.Email,
		CreatedAt:      now,
		UpdatedAt:      now,
		HashedPassword: string(user.Password.hash),
	}

	dbUser, err := m.DB.CreateUser(ctx, *args)

	// Add the values to the user model
	user.ID = dbUser.UserID
	user.Version = int(dbUser.Version)
	user.CreatedAt = dbUser.CreatedAt

	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*db.GetForTokenRow, error) {
	// Calculate the SHA-256 hash of the plaintext token provided by the client.
	// Remember that this returns a byte *array* with length 32, not a slice.
	// tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := &db.GetForTokenParams{
		RefreshToken: tokenPlaintext,
		Expiry:       time.Now(),
	}

	user, err := m.DB.GetForToken(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
