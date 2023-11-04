package token

import (
	"time"
	"errors"

	"github.com/google/uuid"
)


var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	Username string `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	Role string `json:"role"`
	ExpiredAt time.Time `json:"expired_at"`
	ID uuid.UUID `json:"id"`
}

func NewPayload(username string, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssuedAt: time.Now(),
		Role: role,
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
