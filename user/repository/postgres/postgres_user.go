package postgres

import (
	"errors"

	"github.com/tfkhdyt/forum-api-go/domain"

	"gorm.io/gorm"
)

// ==================================

type postgresUserRepository struct {
	conn *gorm.DB
}

// ==================================

func New(conn *gorm.DB) domain.UserRepository {
	return &postgresUserRepository{conn}
}

// ==================================

func (p *postgresUserRepository) Create(
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

func (p *postgresUserRepository) VerifyAvailableUsername(username string) error {
	var result struct {
		username string
	}

	if err := p.conn.Model(&domain.User{}).Where("username = ?", username).First(&result).Error; err == nil {
		return errors.New("Username is not available")
	}

	return nil
}

func (p *postgresUserRepository) FindPasswordByUsername(username string) (string, error) {
	var result struct {
		password string
	}

	if err := p.conn.Model(&domain.User{}).Where("username = ?", username).First(&result).Error; err != nil {
		return "", err
	}

	return result.password, nil
}

func (p *postgresUserRepository) FindIdByUsername(username string) (uint, error) {
	var result struct {
		id uint
	}

	if err := p.conn.Model(&domain.User{}).Where("username = ?", username).First(&result).Error; err != nil {
		return 0, err
	}

	return result.id, nil
}
