package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/philjestin/ranked-talishar/test_util" // Optional: Use a mocking framework for better assertions
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(test_util.RandomString(32))

	require.NoError(t, err)

	user_name := test_util.RandomFirstName()
	// Expiration time passed to NewPayload and is multiplied by 5
	duration := time.Minute * 5

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	tokens, err := maker.CreateToken(user_name, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokens.AccessToken)
	require.NotEmpty(t, tokens.RefreshToken)

	payload, err := maker.VerifyToken(tokens.AccessToken)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, user_name, payload.UserName)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(test_util.RandomString(32))
	require.NoError(t, err)

	tokens, err := maker.CreateToken(test_util.RandomFirstName(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokens.AccessToken)

	payload, err := maker.VerifyToken(tokens.AccessToken)
	require.Error(t, err)
	require.EqualError(t, err, "token has invalid claims: token is expired")

	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(test_util.RandomFirstName(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(test_util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token is unverifiable: error while executing keyfunc: token is unverifiable")
	require.Nil(t, payload)
}
