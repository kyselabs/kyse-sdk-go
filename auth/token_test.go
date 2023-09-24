package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestIsExpiredWithValidToken(t *testing.T) {
	expirationTime := time.Now().Add(time.Hour)
	token := createValidToken(expirationTime)

	expired, err := IsExpired(token)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if expired {
		t.Error("Expected token to be valid, but it's expired.")
	}
}

func createValidToken(expirationTime time.Time) string {
	claims := jwt.MapClaims{
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	signedToken, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	return signedToken
}
