package sqlite

import (
	"github.com/tfkhdyt/forum-api-go/domain"
	"gorm.io/gorm"
)

type SqliteUserRepository struct {
	Conn *gorm.DB
}

func (s *SqliteUserRepository) New(conn *gorm.DB) domain.UserRepository {
	return &SqliteUserRepository{conn}
}

func (s *SqliteUserRepository) Create(
	createUserDto domain.CreateUserDto,
) (domain.CreatedUserDto, error) {
	user := domain.User{
		Username: createUserDto.Username,
		Password: createUserDto.Password,
		FullName: createUserDto.FullName,
	}

	if err := s.Conn.Create(&user).Error; err != nil {
		return domain.CreatedUserDto{}, err
	}

	return domain.CreatedUserDto{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
	}, nil
}

func (s *SqliteUserRepository) VerifyAvailableUsername(username string) error {
	var user domain.User

	if err := s.Conn.Select("username").Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *SqliteUserRepository) FindPasswordByUsername(username string) (string, error) {
	var user domain.User

	if err := s.Conn.Select("password").Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}

func (s *SqliteUserRepository) FindIdByUsername(username string) (uint, error) {
	var user domain.User

	if err := s.Conn.Select("ID").Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
