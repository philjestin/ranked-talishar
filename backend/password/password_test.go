package password

import (
	"testing"

	"github.com/philjestin/ranked-talishar/test_util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
		password := test_util.RandomString(6)

		hashedPassword, err := HashedPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		err = CheckPassword(password, hashedPassword)
		require.NoError(t, err)

		wrongPassword := test_util.RandomString(6)
		err = CheckPassword(wrongPassword, hashedPassword)
		require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}