package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type CreateUserUseCase struct {
	userService services.UserService
}

func NewCreateUserUseCase(us services.UserService) CreateUserUseCase {
	return CreateUserUseCase{userService: us}
}

func (cuuc *CreateUserUseCase) Execute(dto dtos.CreateUserDTO) (*entities.UserEntity, error) {
	userToCreate := entities.UserEntity{
		OrganizationId: nil,
		PositionId:     nil,
		Role:           dto.Role,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
	}
	user, err := cuuc.userService.CreateUser(userToCreate)

	return user, err
}
