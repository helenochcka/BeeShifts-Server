package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type AttachUserUseCase struct {
	userService services.UserService
}

func NewAttachUserUseCase(us services.UserService) AttachUserUseCase {
	return AttachUserUseCase{userService: us}
}

func (auuc *AttachUserUseCase) Execute(dto dtos.AttachUserDTO) (*entities.UserEntity, error) {
	usersFilterDTO := dtos.UsersFilterDTO{Ids: []int{dto.Id}}
	user, err := auuc.userService.GetUser(usersFilterDTO)

	if err != nil {
		return nil, err
	}

	userToUpdate := entities.UserEntity{
		Id:             dto.Id,
		OrganizationId: dto.OrganizationId,
		PositionId:     dto.PositionId,
		Role:           user.Role,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Password:       user.Password,
	}
	attachedUser, err := auuc.userService.UpdateUser(userToUpdate)

	return attachedUser, err
}
