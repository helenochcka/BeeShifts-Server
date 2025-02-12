package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type AttachUserUseCase struct {
	userService     services.UserService
	positionService positionsServices.PositionService
}

func NewAttachUserUseCase(us services.UserService, ps positionsServices.PositionService) AttachUserUseCase {
	return AttachUserUseCase{userService: us, positionService: ps}
}

func (auuc *AttachUserUseCase) Execute(managerId int, dto users.AttachDTO) (*users.Entity, error) {
	filter := users.FilterDTO{Ids: []int{dto.Id}}
	slog.Info("Getting user to attach by filter...", "filter", filter)
	userToUpdate, err := auuc.userService.GetUser(filter)

	if err != nil {
		slog.Error("Error getting user to attach...", "err", err)
		return nil, err
	}

	manager, err := auuc.userService.GetUser(users.FilterDTO{Ids: []int{managerId}})

	if err != nil {
		slog.Error("Error getting manager to attach...", "err", err)
		return nil, err
	}

	slog.Info("Checking organization of user...", "user", userToUpdate)
	if userToUpdate.OrganizationId != nil && *userToUpdate.OrganizationId != *manager.OrganizationId {
		slog.Error("Conflicting by user's organization...", "organization", *userToUpdate.OrganizationId)
		return nil, users.EmployeeAlreadyAttached
	}

	positionFilter := positions.FilterDTO{
		Ids:        []int{dto.PositionId},
		ManagerIds: []int{managerId},
	}

	slog.Info("Checking position by filter...", "filter", positionFilter)
	_, err = auuc.positionService.GetPosition(positionFilter)

	if err != nil {
		slog.Error("Error getting position to attach...", "error", err)
		return nil, err
	}

	slog.Info("Changing organization and position of user by dto...", "dto", dto)
	userToUpdate.OrganizationId = manager.OrganizationId
	userToUpdate.PositionId = &dto.PositionId

	slog.Info("Updating user to attach...", "userToUpdate", userToUpdate)
	attachedUser, err := auuc.userService.UpdateUser(*userToUpdate)

	if err != nil {
		slog.Error("Error updating user to attach...", "err", err)
		return nil, err
	}

	return attachedUser, nil
}
