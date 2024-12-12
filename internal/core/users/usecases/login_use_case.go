package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
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

	slog.Info("Getting user for login by filter...", "filter", filter)
	user, err := luc.userService.FindUser(filter)

	if err != nil {
		slog.Error("Error getting user...", "error", err)
		return nil, err
	}

	slog.Info("Checking credentials for user...", "user", user, "creds", creds)
	if user == nil || user.Password != creds.Password {
		slog.Error("Error getting user...", "user", user, "creds", creds)
		return nil, users.IncorrectCredentials
	}

	slog.Info("Generating access token by user id...", "userId", user.Id)
	token, err := luc.authService.GenerateToken(user.Id)

	if err != nil {
		slog.Error("Error generating access token...", "err", err)
		return nil, err
	}
	return &users.LoginResponseDTO{AccessToken: token}, nil
}
