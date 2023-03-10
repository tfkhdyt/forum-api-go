package domain

type CryptoService interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hashedPassword string) error
}
