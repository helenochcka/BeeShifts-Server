package usecases

import (
	"BeeShifts-Server/internal/core/users/services"
)

type AuthenticateUseCase struct {
	authService services.AuthService
}

func NewAuthenticateUseCase(as services.AuthService) AuthenticateUseCase {
	return AuthenticateUseCase{authService: as}
}

func (auc *AuthenticateUseCase) Execute(token string) (*int, error) {
	tokenPayload, err := auc.authService.PayloadFromToken(token)

	if err != nil {
		return nil, err
	}

	return &tokenPayload.Id, nil
}
