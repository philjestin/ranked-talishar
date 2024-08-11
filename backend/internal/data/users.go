package data

import (
	"context"
	"database/sql"
	"time"

	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

type UserModel struct {
	db *db.Queries
}

// Declare a new AnonymousUser variable.
var AnonymousUser = &User{}

type User struct {
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	Password  string    `json:"password"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Check if a User instance is the AnonymousUser.
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

	_, err := m.db.CreateUser(ctx, *args)
	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) GetForToken(tokenPlaintext string) (*db.GetForTokenRow, error) {
	// Calculate the SHA-256 hash of the plaintext token provided by the client.
	// Remember that this returns a byte *array* with length 32, not a slice.
	// tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := &db.GetForTokenParams{
		RefreshToken: tokenPlaintext,
		Expiry:       time.Now(),
	}

	user, err := m.db.GetForToken(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
