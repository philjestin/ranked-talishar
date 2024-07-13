package password

import (
	"testing"

	"github.com/philjestin/ranked-talishar/test_util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
		password := test_util.RandomString(6)

		hashedPassword1, err := HashedPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword1)

		err = CheckPassword(password, hashedPassword1)
		require.NoError(t, err)

		wrongPassword := test_util.RandomString(6)
		err = CheckPassword(wrongPassword, hashedPassword1)
		require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

		hashedPassword2, err := HashedPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword2)
		require.NotEqual(t, hashedPassword1, hashedPassword2)
}