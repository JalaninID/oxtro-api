package pkg

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// encrypt password
func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
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

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[seededRand.Intn(len(letter))]
	}
	return string(b)
}
