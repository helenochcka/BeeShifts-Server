package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
	"errors"
)

type AuthnUseCase struct {
	userService services.UserService
}

func NewAuthnUseCase(us services.UserService) AuthnUseCase {
	return AuthnUseCase{userService: us}
}

func (auc *AuthnUseCase) Execute(role string, id int) error {
	dto := dtos.UsersFilterDTO{
		Ids: []int{id},
	}
	user, err := auc.userService.GetUser(dto)

	if err != nil {
		return err
	}

	if user.Role != role {
		return errors.New("access is denied") //401
	}

	return nil
}
