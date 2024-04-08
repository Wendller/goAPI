package cryptography

import (
	"log"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

type JWTEncrypter struct{}

func NewJWTEncrypter() *JWTEncrypter {
	return &JWTEncrypter{}
}

func (e *JWTEncrypter) Encrypt(payload map[string]interface{}) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	JWTAuthConfig := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

	_, tokenString, _ := JWTAuthConfig.Encode(map[string]interface{}{
		"sub": payload["userID"],
		"exp": time.Now().Add(time.Second * time.Duration(300)).Unix(),
	})

	return tokenString
}
