package postgres

import (
	"errors"

	"github.com/tfkhdyt/forum-api-go/domain"
	"gorm.io/gorm"
)

type postgresAuthRepository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) domain.AuthRepository {
	return &postgresAuthRepository{conn}
}

func (p *postgresAuthRepository) CreateToken(token string) error {
	auth := domain.Auth{
		RefreshToken: token,
	}

	if err := p.conn.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}

func (p *postgresAuthRepository) CheckTokenAvailability(token string) error {
	var auth domain.Auth

	if err := p.conn.Where("refresh_token = ?", token).First(&auth).Error; err != nil {
		return errors.New("Refresh token is not found in database")
	}

	return nil
}

func (p *postgresAuthRepository) DeleteToken(token string) error {
	if err := p.conn.Where("refresh_token = ?", token).Delete(&domain.Auth{}).Error; err != nil {
		return err
	}

	return nil
}
