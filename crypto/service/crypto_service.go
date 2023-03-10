package service

import (
	"github.com/tfkhdyt/forum-api-go/domain"
	"golang.org/x/crypto/bcrypt"
)

type bcryptCryptoService struct{}

func New() domain.CryptoService {
	return &bcryptCryptoService{}
}

func (c *bcryptCryptoService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hashed), err
}

func (c *bcryptCryptoService) ComparePassword(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}
