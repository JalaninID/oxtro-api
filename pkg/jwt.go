package pkg

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID  string `json:"id"`
	Exp string `json:"exp"`
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]any, expired int) (string, error) {
	duration, _ := strconv.Atoi(os.Getenv("JWT_TIME_DURATION"))
	if expired > 0 {
		duration = expired
	}

	drt := time.Minute * time.Duration(duration)
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(drt).Unix()

	for i, v := range Data {
		claims[i] = v
	}
	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func VerifyTokenHeader(requestToken string) (MetaToken, error) {

	token, err := jwt.Parse((requestToken), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return MetaToken{}, err
	}
	claimToken := DecodeToken(token)
	return claimToken.Claims, nil
}

func VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, err := json.Marshal(&accessToken)
	if err != nil {
		return token
	}
	err = json.Unmarshal(stringify, &token)
	if err != nil {
		return token
	}
	return token
}
