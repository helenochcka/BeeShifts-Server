package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type DetachUserUseCase struct {
	userService services.UserService
}

func NewDetachUserUseCase(us services.UserService) DetachUserUseCase {
	return DetachUserUseCase{userService: us}
}

func (duuc *DetachUserUseCase) Execute(managerId int, dto users.DetachDTO) (*users.Entity, error) {
	slog.Info("Getting user to detach by id...", "id", dto.Id)
	userToDetach, err := duuc.userService.GetUser(users.FilterDTO{Ids: []int{dto.Id}})

	if err != nil {
		slog.Error("Error getting user to detach...", "err", err)
		return nil, err
	}

	manager, err := duuc.userService.GetUser(users.FilterDTO{Ids: []int{managerId}})

	if err != nil {
		slog.Error("Error getting manager to detach...", "err", err)
		return nil, err
	}

	slog.Info("Checking organization of user...", "user", userToDetach)
	if userToDetach.OrganizationId == nil || *userToDetach.OrganizationId != *manager.OrganizationId {
		slog.Error("Conflicting by user's organization...", "organization", *userToDetach.OrganizationId)
		return nil, users.EmployeeNotAttached
	}

	userToDetach.OrganizationId = nil
	userToDetach.PositionId = nil

	slog.Info("Updating user to detach...", "userToDetach", userToDetach)
	detachedUser, err := duuc.userService.UpdateUser(*userToDetach)

	if err != nil {
		slog.Error("Error updating user to detach...", "err", err)
		return nil, err
	}

	return detachedUser, nil
}
