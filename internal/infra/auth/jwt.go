package auth

import (
	"log"
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

type JWTAuthConfig struct {
	JWT            *jwtauth.JWTAuth
	JWT_SECRET     string
	JWT_EXPIRES_IN string
}

func NewJWTAuthConfig() *JWTAuthConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpiresIn := os.Getenv("JWT_EXPIRES_IN")

	jwtAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

	return &JWTAuthConfig{
		JWT:            jwtAuth,
		JWT_SECRET:     jwtSecret,
		JWT_EXPIRES_IN: jwtExpiresIn,
	}
}
