package cryptography

import (
	"strconv"
	"time"

	"github.com/Wendller/goexpert/goAPI/internal/infra/auth"
)

type JWTEncrypter struct{}

func NewJWTEncrypter() *JWTEncrypter {
	return &JWTEncrypter{}
}

func (e *JWTEncrypter) Encrypt(payload map[string]interface{}) string {
	JWTAuthConfig := auth.NewJWTAuthConfig()
	tokenExpires, _ := strconv.Atoi(JWTAuthConfig.JWT_EXPIRES_IN)

	_, tokenString, _ := JWTAuthConfig.JWT.Encode(map[string]interface{}{
		"sub": payload["userID"],
		"exp": time.Now().Add(time.Second * time.Duration(tokenExpires)).Unix(),
	})

	return tokenString
}
