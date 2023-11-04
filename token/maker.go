package token

import "time"


type Maker interface {
	CreateToken(string, string, time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
