package auth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	Anonymous int64 = 0
	Admin     int64 = 9999
)

var (
	tokenSecretKey = "jwttokensecretkey"
)

// SetTokenSecretKey replaces token secret key with k.
func SetTokenSecretKey(k string) {
	tokenSecretKey = k
}

// AuthInfo
type AuthInfo struct {
	jwt.StandardClaims
	Id   string `json:"id,omitempty"`
	Role int64  `json:"role,omitempty"`
}

// IsExpired return true if the Info is expired.
func (i *AuthInfo) IsExpired() bool {
	if i.ExpiresAt < time.Now().Unix() {
		return true
	}
	return false
}

// EncryptToken returns a token string using Info given by InfoProvider ip.
func EncryptToken(ai AuthInfo) (string, error) {
	if ai.Id == "" {
		return "", errors.New("Id was empty.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ai)
	encryptedToken, err := token.SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return encryptedToken, nil
}

// DecryptToken decrypts the token string and returns a ptr of Info.
func DecryptToken(encryptedToken string) (*AuthInfo, error) {
	ai := AuthInfo{}
	// Parse, validate and return a token.
	_, err := jwt.ParseWithClaims(encryptedToken, &ai, func(token *jwt.Token) (interface{}, error) {
		// Prevents a known exploit
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return tokenSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return &ai, nil
}
