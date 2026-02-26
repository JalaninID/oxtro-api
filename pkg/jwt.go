package pkg

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MetaToken struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Exp       int64  `json:"exp"`
	TokenType string `json:"token_type"`
}

type AccessToken struct {
	Claims MetaToken
}

type Claims struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func Sign(data map[string]any, expired int) (string, error) {
	duration, _ := strconv.Atoi(os.Getenv("JWT_TIME_DURATION"))
	if expired > 0 {
		duration = expired
	}

	ttl := time.Minute * time.Duration(duration)
	expiresAt := time.Now().Add(ttl)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	if id, ok := data["id"].(string); ok {
		claims.ID = id
	}
	if sessionID, ok := data["session_id"].(string); ok {
		claims.SessionID = sessionID
	}
	if tokenType, ok := data["token_type"].(string); ok {
		claims.TokenType = tokenType
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func VerifyTokenHeader(requestToken string) (MetaToken, error) {
	token, err := VerifyToken(requestToken)
	if err != nil {
		return MetaToken{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return MetaToken{}, errors.New("invalid token claims")
	}

	return MetaToken{
		ID:        claims.ID,
		SessionID: claims.SessionID,
		TokenType: claims.TokenType,
		Exp:       claims.ExpiresAt.Unix(),
	}, nil
}

func VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	claims, ok := accessToken.Claims.(*Claims)
	if !ok || claims.ExpiresAt == nil {
		return AccessToken{}
	}
	return AccessToken{
		Claims: MetaToken{
			ID:        claims.ID,
			SessionID: claims.SessionID,
			Exp:       claims.ExpiresAt.Unix(),
			TokenType: claims.TokenType,
		},
	}
}
