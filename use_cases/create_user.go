package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
)

type CreateUserUseCase struct {
	userService services.UserService
}

func NewCreateUserUseCase(us services.UserService) CreateUserUseCase {
	return CreateUserUseCase{userService: us}
}

func (cuuc *CreateUserUseCase) Execute(dto dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	user, err := cuuc.userService.CreateUser(dto)

	return user, err
}
