package service

import "github.com/tfkhdyt/forum-api-go/domain"

type authService struct {
	authRepo      domain.AuthRepository
	userRepo      domain.UserRepository
	cryptoService domain.CryptoService
	tokenService  domain.TokenService
}

func New(
	authRepo domain.AuthRepository,
	userRepo domain.UserRepository,
	cryptoService domain.CryptoService,
	tokenService domain.TokenService,
) domain.AuthService {
	return &authService{authRepo, userRepo, cryptoService, tokenService}
}

func (a *authService) Login(loginDto domain.LoginDto) (domain.Credentials, error) {
	username, password := loginDto.Username, loginDto.Password
	hashedPassword, err := a.userRepo.FindPasswordByUsername(username)
	if err != nil {
		return domain.Credentials{}, err
	}

	if err := a.cryptoService.ComparePassword(password, hashedPassword); err != nil {
		return domain.Credentials{}, err
	}

	id, err := a.userRepo.FindIdByUsername(username)
	if err != nil {
		return domain.Credentials{}, err
	}

	accessToken, err := a.tokenService.CreateAccessToken(id)
	if err != nil {
		return domain.Credentials{}, err
	}

	refreshToken, err := a.tokenService.CreateRefreshToken(id)
	if err != nil {
		return domain.Credentials{}, err
	}

	if err := a.authRepo.CreateToken(refreshToken); err != nil {
		return domain.Credentials{}, err
	}

	return domain.Credentials{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a *authService) Logout(deleteAuthDto domain.LogoutRefreshDto) error {
	token := deleteAuthDto.RefreshToken

	if err := a.authRepo.CheckTokenAvailability(token); err != nil {
		if err := a.authRepo.DeleteToken(token); err != nil {
			return err
		}
	}

	return nil
}

func (a *authService) RefreshToken(refreshTokenDto domain.LogoutRefreshDto) (string, error) {
	refreshToken := refreshTokenDto.RefreshToken

	userId, err := a.tokenService.DecodePayload(refreshToken)
	if err != nil {
		return "", err
	}

	if err := a.authRepo.CheckTokenAvailability(refreshToken); err != nil {
		return "", err
	}

	accessToken, err := a.tokenService.CreateAccessToken(userId)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
