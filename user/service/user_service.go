package service

import "github.com/tfkhdyt/forum-api-go/domain"

type UserService struct {
	UserRepo domain.UserRepository
}

func (u *UserService) New(userRepo domain.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (u *UserService) CreateUser(
	createUserDto domain.CreateUserDto,
) (domain.CreatedUserDto, error) {
	if err := u.UserRepo.VerifyAvailableUsername(createUserDto.Username); err != nil {
		return domain.CreatedUserDto{}, err
	}

	hashedPassword, err := hashPassword(createUserDto.Password)
	if err != nil {
		return domain.CreatedUserDto{}, err
	}

	createUserDto.Password = hashedPassword

	createdUser, err := u.UserRepo.CreateUser(createUserDto)
	if err != nil {
		return domain.CreatedUserDto{}, err
	}

	return createdUser, nil
}
