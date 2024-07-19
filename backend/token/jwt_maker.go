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

type CreateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(user_name string, duration time.Duration) (CreateTokenResponse, error) {
	payload, err := NewPayload(user_name, duration)
	if err != nil {
		response := CreateTokenResponse{
			AccessToken:  "",
			RefreshToken: "",
		}
		return response, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		response := CreateTokenResponse{
			AccessToken:  "",
			RefreshToken: "",
		}
		return response, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refreshToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		response := CreateTokenResponse{
			AccessToken:  "",
			RefreshToken: "",
		}
		return response, err
	}

	response := CreateTokenResponse{
		AccessToken:  token,
		RefreshToken: rt,
	}

	return response, err
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

// this is the same as above, but i might want to change it later
func (maker *JWTMaker) VerifyRefreshToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	payload := &Payload{}
	fmt.Println("payload", payload)
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
