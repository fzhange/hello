package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
beacuse we use symmetric key algorithm to sign the tokens
so this struct's secretKey field will strore secret key.
*/
type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	/**
	What is a key function? Well,
	basically, it’s a function that receives the parsed but unverified token.
	You should verify its header
	to make sure that the signing algorithm matches with what you normally use to sign the tokens.

	Then if it matches, you return the key so that jwt-go can use it to verify the token
	 **/

	/**
	In the key function,
	we can get its signing algorithm via the token.Method field.
	Note that its type is a SigningMethod, which is just an interface.
	So we have to try converting it to a specific implementation.

	In our case, we convert it to SigningMethodHMAC because we’re using HS256,
	which is an instance of the SigningMethodHMAC struct.
	*/

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			//the algorithm of the token doesn’t match with our signing algorithm
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		// there might be 2 different scenarios:
		// either the token is invalid or it is expired.
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
