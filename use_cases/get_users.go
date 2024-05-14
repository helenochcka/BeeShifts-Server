package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
)

type GetUserUseCase struct {
	userService services.UserService
}

func NewGetUserUseCase(us services.UserService) GetUserUseCase {
	return GetUserUseCase{userService: us}
}

func (guuc *GetUserUseCase) Execute(dto dtos.UsersFilterDTO) ([]dtos.UserDTO, error) {
	users, err := guuc.userService.GetUsers(dto)

	return users, err
}
