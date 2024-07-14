package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrorInvalidToken = errors.New("token is unverifiable")
	ErrExpiredToken   = errors.New("token has expired")
)

type Payload struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func NewPayload(user_name string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(duration)

	payload := &Payload{
		UserName: user_name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
