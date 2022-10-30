package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

/*
Normally these 3 fields should be enough. However,
if we want to have a mechanism to invalidate some specific tokens in case they are leaked,
we need to add an ID field to uniquely identify each token
**/
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"` //identify the token owner.
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload must implement Valid method.
// beacuse jwt-go package need this to check token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
