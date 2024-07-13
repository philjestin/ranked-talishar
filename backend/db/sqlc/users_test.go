package db

import (
	"context"
	"testing"

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
