package token

import "time"

type Maker interface {
	CreateToken(user_name string, duration time.Duration) (CreateTokenResponse, error)

	VerifyToken(token string) (*Payload, error)
	VerifyRefreshToken(token string) (*Payload, error)
}
