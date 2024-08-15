package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

type UserModel struct {
	DB *db.Queries
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

var AnonymousUser = &User{}

// type password struct {
// 	plaintext *string
// 	hash      []byte
// }

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (m UserModel) Insert(user db.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now()
	args := &db.CreateUserParams{
		UserName:       user.UserName,
		UserEmail:      user.UserEmail,
		CreatedAt:      now,
		UpdatedAt:      now,
		HashedPassword: user.HashedPassword,
	}

	_, err := m.DB.CreateUser(ctx, *args)
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
