package service

import "github.com/tfkhdyt/forum-api-go/domain"

// ======================

type UserService struct {
	userRepo domain.UserRepository
}

// ======================

func New(userRepo domain.UserRepository) domain.UserService {
	return &UserService{userRepo}
}

// ======================

func (u *UserService) Create(
	createUserDto domain.CreateUserDto,
) (domain.CreatedUserDto, error) {
	if err := u.userRepo.VerifyAvailableUsername(createUserDto.Username); err != nil {
		return domain.CreatedUserDto{}, err
	}

	hashedPassword, err := hashPassword(createUserDto.Password)
	if err != nil {
		return domain.CreatedUserDto{}, err
	}

	createUserDto.Password = hashedPassword

	createdUser, err := u.userRepo.Create(createUserDto)
	if err != nil {
		return domain.CreatedUserDto{}, err
	}

	return createdUser, nil
}
