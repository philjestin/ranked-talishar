package token

import "time"

type Maker interface {
	CreateToken(user_name string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
