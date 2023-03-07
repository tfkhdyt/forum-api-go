package postgres

import (
	"errors"

	"github.com/tfkhdyt/forum-api-go/domain"

	"gorm.io/gorm"
)

// ==================================

type PostgresUserRepository struct {
	conn *gorm.DB
}

// ==================================

func New(conn *gorm.DB) domain.UserRepository {
	return &PostgresUserRepository{conn}
}

// ==================================

func (p *PostgresUserRepository) Create(
	createUserDto domain.CreateUserDto,
) (domain.CreatedUserDto, error) {
	user := domain.User{
		Username: createUserDto.Username,
		Password: createUserDto.Password,
		FullName: createUserDto.FullName,
	}

	if err := p.conn.Create(&user).Error; err != nil {
		return domain.CreatedUserDto{}, err
	}

	return domain.CreatedUserDto{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
	}, nil
}

func (p *PostgresUserRepository) VerifyAvailableUsername(username string) error {
	var result struct {
		username string
	}

	if err := p.conn.Model(&domain.User{}).Where("username = ?", username).First(&result).Error; err == nil {
		return errors.New("Username is not available")
	}

	return nil
}

func (p *PostgresUserRepository) FindPasswordByUsername(username string) (string, error) {
	var user domain.User

	if err := p.conn.Select("password").Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}

func (p *PostgresUserRepository) FindIdByUsername(username string) (uint, error) {
	var user domain.User

	if err := p.conn.Select("ID").Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
