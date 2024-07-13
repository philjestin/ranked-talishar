package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/philjestin/ranked-talishar/password"
	"github.com/philjestin/ranked-talishar/test_util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := password.HashedPassword(test_util.RandomString(6))

	require.NoError(t, err)

	arg := CreateUserParams{
		UserName:       test_util.RandomFirstName(),
		HashedPassword: hashedPassword,
		UserEmail:      test_util.RandomString(12),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.UserEmail, user.UserEmail)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.UserName, user2.UserName)
	require.Equal(t, user1.UserEmail, user2.UserEmail)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		UserID:   user1.UserID,
		UserName: sql.NullString{String: test_util.RandomFirstName()},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	// require.Equal(t, arg.UserName.String, user2.UserName)
	require.Equal(t, user1.UserEmail, user2.UserEmail)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

	require.NotZero(t, user2.UpdatedAt)
}
