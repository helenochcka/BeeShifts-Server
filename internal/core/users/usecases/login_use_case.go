package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"errors"
)

type LoginUseCase struct {
	userService services.UserService
	authService services.AuthService
}

func NewLoginUseCase(us services.UserService, as services.AuthService) LoginUseCase {
	return LoginUseCase{userService: us, authService: as}
}

func (luc *LoginUseCase) Execute(creds users.CredsDTO) (*users.LoginResponseDTO, error) {
	filter := users.FilterDTO{
		Emails: []string{creds.Email},
	}

	user, err := luc.userService.GetUser(filter)

	if err != nil {
		if errors.Is(err, users.UserNotFound) {
			err = users.IncorrectCredentials
		}
		return nil, err
	}

	if user.Password != creds.Password {
		return nil, users.IncorrectCredentials
	}

	token, err := luc.authService.GenerateToken(user.Id)

	if err != nil {
		return nil, err
	}
	return &users.LoginResponseDTO{AccessToken: token}, nil
}
