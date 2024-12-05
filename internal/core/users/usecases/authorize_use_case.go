package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
)

type AuthorizeUseCase struct {
	userService services.UserService
}

func NewAuthorizeUseCase(us services.UserService) AuthorizeUseCase {
	return AuthorizeUseCase{userService: us}
}

func (auc *AuthorizeUseCase) Execute(role string, id int) error {
	filter := users.FilterDTO{
		Ids: []int{id},
	}
	user, err := auc.userService.GetUser(filter)

	if err != nil {
		return err
	}

	if user.Role != role {
		return users.InsufficientRights //401
	}

	return nil
}
