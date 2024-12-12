package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type AttachUserUseCase struct {
	userService services.UserService
}

func NewAttachUserUseCase(us services.UserService) AttachUserUseCase {
	return AttachUserUseCase{userService: us}
}

func (auuc *AttachUserUseCase) Execute(dto users.AttachDTO) (*users.Entity, error) {
	filter := users.FilterDTO{Ids: []int{dto.Id}}
	slog.Info("Getting user to attach by filter...", "filter", filter)
	userToUpdate, err := auuc.userService.GetUser(filter)

	if err != nil {
		slog.Error("Error getting user to attach...", "err", err)
		return nil, err
	}

	slog.Info("Changing organization and position of user by dto...", "dto", dto)
	userToUpdate.OrganizationId = dto.OrganizationId
	userToUpdate.PositionId = dto.PositionId

	slog.Info("Updating user to attach...", "userToUpdate", userToUpdate)
	attachedUser, err := auuc.userService.UpdateUser(*userToUpdate)

	if err != nil {
		slog.Error("Error updating user to attach...", "err", err)
		return nil, err
	}

	return attachedUser, nil
}
