package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
	"errors"
)

type CreateUserUseCase struct {
	userService services.UserService
}

func NewCreateUserUseCase(us services.UserService) CreateUserUseCase {
	return CreateUserUseCase{userService: us}
}

func (cuuc *CreateUserUseCase) Execute(dto dtos.CreateUserDTO) (*entities.UserEntity, error) {
	userFilter := dtos.UsersFilterDTO{
		Emails: []string{dto.Email},
	}
	user, err := cuuc.userService.GetUser(userFilter)
	if user != nil {
		return nil, errors.New("emails already used")
	}

	userToCreate := entities.UserEntity{
		OrganizationId: nil,
		PositionId:     nil,
		Role:           dto.Role,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
	}
	createdUser, err := cuuc.userService.CreateUser(userToCreate)

	return createdUser, err
}
