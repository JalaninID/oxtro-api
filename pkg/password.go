package pkg

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// encrypt password
func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	if configured := os.Getenv("BCRYPT_COST"); configured != "" {
		if parsedCost, err := strconv.Atoi(configured); err == nil && parsedCost >= bcrypt.MinCost && parsedCost <= bcrypt.MaxCost {
			cost = parsedCost
		}
	}
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	passwordHash := string(bytePassword)

	return passwordHash, nil
}

// compare password
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func ValidatePassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	// var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		// case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
		// 	specialCharPresent = true
		// 	passLen++
		case ch == ' ':
			passLen++
		}
	}
	if !lowercasePresent {
		return errors.New("Invalid password at least 1 lowercase letter")
	}
	if !uppercasePresent {
		return errors.New("Invalid password at least 1 uppercase letter")
	}
	if !numberPresent {
		return errors.New("Invalid password at least 1 numeric letter")
	}
	// if !specialCharPresent {
	// appendError("special character missing")
	// }
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		return fmt.Errorf("password length must be between %d to %d characters long", minPassLength, maxPassLength)

	}
	return nil
}

func GenerateRandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letter))))
		if err != nil {
			seeded := time.Now().UnixNano() % int64(len(letter))
			b[i] = letter[seeded]
			continue
		}
		b[i] = letter[idx.Int64()]
	}
	return string(b)
}
