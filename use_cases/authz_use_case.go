package use_cases

import (
	"BeeShifts-Server/services"
	"github.com/golang-jwt/jwt"
)

type AuthzUseCase struct {
	authzService services.AuthzService
}

func NewAuthzUseCase(as services.AuthzService) AuthzUseCase {
	return AuthzUseCase{authzService: as}
}

func (auc *AuthzUseCase) Execute(token string) (jwt.MapClaims, error) {
	payload, err := auc.authzService.PayloadFromToken(token)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
