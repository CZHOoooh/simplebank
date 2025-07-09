package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

type JwtClaims struct {
	ID       uuid.UUID
	Username string
	jwt.RegisteredClaims
}

// CreateToken creates a new token for a special username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *PayloadToClaims(payload))
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	jwtClaims, ok := jwtToken.Claims.(*JwtClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return ClaimsToPayload(jwtClaims), nil
}

func PayloadToClaims(payload *Payload) *JwtClaims {
	return &JwtClaims{
		ID:       payload.ID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
	}
}

func ClaimsToPayload(jwtClaims *JwtClaims) *Payload {
	return &Payload{
		ID:        jwtClaims.ID,
		Username:  jwtClaims.Username,
		IssuedAt:  jwtClaims.RegisteredClaims.IssuedAt.Time,
		ExpiredAt: jwtClaims.RegisteredClaims.ExpiresAt.Time,
	}
}
