package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type UpdateUserUseCase struct {
	userService services.UserService
}

func NewUpdateUserUseCase(us services.UserService) UpdateUserUseCase {
	return UpdateUserUseCase{userService: us}
}

func (uuuc *UpdateUserUseCase) Execute(id int, dto users.UpdateSelfDTO) (*users.Entity, error) {
	filter := users.FilterDTO{Ids: []int{id}}
	slog.Info("Getting user to update by filter...", "filter", filter)
	userToUpdate, err := uuuc.userService.GetUser(filter)

	if err != nil {
		slog.Error("Error getting user...", "err", err)
		return nil, err
	}

	filter = users.FilterDTO{Emails: []string{dto.Email}}
	slog.Info("Looking for conflicting user by filter...", "filter", filter)
	conflictingUser, err := uuuc.userService.FindUser(filter)

	if err != nil {
		slog.Error("Error getting conflicting user...", "err", err)
	}

	slog.Info("Checking if found conflicting user isn't the same as user to update...")
	if conflictingUser != nil && conflictingUser.Id != userToUpdate.Id {
		slog.Error("Conflicting by email user is found...", "email", dto.Email)
		return nil, users.EmailAlreadyUsed
	}

	userToUpdate.FirstName = dto.FirstName
	userToUpdate.LastName = dto.LastName
	userToUpdate.Email = dto.Email
	userToUpdate.Password = dto.Password

	slog.Info("Updating user...", "userToUpdate", userToUpdate)
	updatedUser, err := uuuc.userService.UpdateUser(*userToUpdate)

	if err != nil {
		slog.Error("Error updating user...", "err", err)
		return nil, err
	}

	return updatedUser, nil
}
