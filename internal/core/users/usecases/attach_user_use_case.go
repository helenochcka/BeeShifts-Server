package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
)

type AttachUserUseCase struct {
	userService services.UserService
}

func NewAttachUserUseCase(us services.UserService) AttachUserUseCase {
	return AttachUserUseCase{userService: us}
}

func (auuc *AttachUserUseCase) Execute(dto users.AttachDTO) (*users.Entity, error) {
	filter := users.FilterDTO{Ids: []int{dto.Id}}
	userToUpdate, err := auuc.userService.GetUser(filter)

	if err != nil {
		return nil, err
	}

	userToUpdate.Id = dto.Id
	userToUpdate.OrganizationId = dto.OrganizationId
	userToUpdate.PositionId = dto.PositionId

	attachedUser, err := auuc.userService.UpdateUser(*userToUpdate)

	return attachedUser, err
}
