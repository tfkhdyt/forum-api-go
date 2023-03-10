package config

import (
	"os"
)

func GetJwtSecretKey() []byte {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	return []byte(jwtSecretKey)
}
