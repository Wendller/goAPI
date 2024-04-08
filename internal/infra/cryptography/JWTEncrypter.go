package cryptography

import (
	"time"

	"github.com/go-chi/jwtauth"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type JWTEncrypter struct{}

func NewJWTEncrypter() *JWTEncrypter {
	return &JWTEncrypter{}
}

func (e *JWTEncrypter) Encrypt(payload map[string]interface{}) string {
	JWTAuthConfig := jwtauth.New("HS256", []byte("secret"), nil)

	_, tokenString, _ := JWTAuthConfig.Encode(map[string]interface{}{
		"sub": payload["userID"],
		"exp": time.Now().Add(time.Second * time.Duration(300)).Unix(),
	})

	return tokenString
}
