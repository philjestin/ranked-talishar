package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

type UserModel struct {
	db  *db.Queries
	ctx context.Context
}

func UserModelController(db *db.Queries, ctx context.Context) *UserModel {
	return &UserModel{db, ctx}
}

// Declare a new AnonymousUser variable.
var AnonymousUser = &User{}

type User struct {
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	Password  string    `json:"password"`
	UserId    uuid.UUID `json:"user_id"`
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

func (cc *UserModel) GetForToken(tokenPlaintext string) (*User, error) {
	// Calculate the SHA-256 hash of the plaintext token provided by the client.
	// Remember that this returns a byte *array* with length 32, not a slice.
	// tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// args := &db.GetForTokenParams{
	// 	RefreshToken: tokenPlaintext,
	// 	// Expiry:       time.Now(),
	// }

	fmt.Println("args inside of GetForToken", tokenPlaintext)

	user, err := cc.db.GetForToken(ctx, tokenPlaintext)

	fmt.Println("user", user)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	res := User{
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		CreatedAt: user.CreatedAt,
		UserId:    user.UserID,
		Password:  user.HashedPassword,
	}

	return &res, nil
}
