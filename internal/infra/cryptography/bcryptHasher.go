package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (hasher *BcryptHasher) Hash(plainText string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func (hasher *BcryptHasher) Compare(hash, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))

	return err == nil
}
