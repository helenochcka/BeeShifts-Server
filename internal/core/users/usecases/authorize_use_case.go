package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
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
	slog.Info("Getting user to authorize by filter...", "filter", filter)
	user, err := auc.userService.GetUser(filter)

	if err != nil {
		slog.Error("Error getting user to authorize...", "err", err)
		return err
	}

	slog.Error("Checking if user has sufficient role...", "user", user, "role", role)
	if user.Role != role {
		slog.Error("User doesn't have sufficient rights...", "user", user, "role", role)
		return users.InsufficientRights
	}

	return nil
}
