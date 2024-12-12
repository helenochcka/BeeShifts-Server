package usecases

import (
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type AuthenticateUseCase struct {
	authService services.AuthService
}

func NewAuthenticateUseCase(as services.AuthService) AuthenticateUseCase {
	return AuthenticateUseCase{authService: as}
}

func (auc *AuthenticateUseCase) Execute(token string) (*int, error) {
	slog.Info("Getting payload from token...", "token", token)
	tokenPayload, err := auc.authService.PayloadFromToken(token)

	if err != nil {
		slog.Error("Error getting payload from token...", "err", err)
		return nil, err
	}

	return &tokenPayload.Id, nil
}
