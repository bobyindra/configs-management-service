package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/golang-jwt/jwt/v5"
)

type TokenParam struct {
	Subject string
}

type AdditionalClaim struct {
	Role string
}

type ConfigsJWTClaim struct {
	jwt.RegisteredClaims
	AdditionalClaim
}

var (
	ErrInvalidToken = func(errorStr string) error {
		return entity.NewError("INVALID_TOKEN", fmt.Sprintf("invalid token: %s", errorStr), http.StatusUnauthorized)
	}
	ErrGenerateTokenFailed = func(errorStr string) error {
		return entity.NewError("FAILED_GENERATE_TOKEN", fmt.Sprintf("generate token failed: %s", errorStr), http.StatusInternalServerError)
	}
	ErrExpiredToken = entity.NewError("EXPIRED_TOKEN", "token has been expired", http.StatusUnauthorized)
)

type auth struct {
	secret             []byte
	expirationDuration time.Duration
}

type Auth interface {
	GenerateToken(param *TokenParam, additionalClaim *AdditionalClaim) (string, error)
	ValidateClaim(ctx context.Context, r *http.Request) (*ConfigsJWTClaim, error)
}

func NewAuth(secret []byte, expirationDuration time.Duration) *auth {
	return &auth{
		secret:             secret,
		expirationDuration: expirationDuration,
	}
}

func (a *auth) GenerateToken(param *TokenParam, additionalClaim *AdditionalClaim) (string, error) {
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(a.expirationDuration)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &ConfigsJWTClaim{
		jwt.RegisteredClaims{
			Subject:   param.Subject,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
		*additionalClaim,
	}

	signed, err := token.SignedString(a.secret)
	if err != nil {
		return "", ErrGenerateTokenFailed(err.Error())
	}

	return signed, nil
}

func (a *auth) ValidateClaim(ctx context.Context, r *http.Request) (*ConfigsJWTClaim, error) {
	token, err := a.getToken(r)
	if err != nil {
		return nil, ErrInvalidToken(err.Error())
	}

	_, claims, err := a.parseToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (a *auth) getToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	if parts[1] == "" {
		return "", errors.New("empty token")
	}

	return parts[1], nil
}

func (a *auth) parseToken(tokenString string) (*jwt.Token, *ConfigsJWTClaim, error) {
	claims := &ConfigsJWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return a.secret, nil
	})

	if err != nil {
		if a.isTokenExpiredError(err) {
			return nil, nil, ErrExpiredToken
		}

		return nil, nil, ErrInvalidToken(err.Error())
	}

	return token, claims, nil
}

func (a *auth) isTokenExpiredError(err error) bool {
	return strings.Contains(err.Error(), jwt.ErrTokenExpired.Error())
}
