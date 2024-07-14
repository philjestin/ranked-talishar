package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(user_name string, duration time.Duration) (string, error) {
	payload, err := NewPayload(user_name, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	payload := &Payload{}
	tokenClaims, err := jwt.ParseWithClaims(token, payload, keyFunc)
	if err != nil {
		// Handle the error appropriately
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrorInvalidToken
		}
		return nil, err
	}

	if !tokenClaims.Valid { // Check for valid claims
		return nil, ErrExpiredToken
	}

	return payload, nil
}
