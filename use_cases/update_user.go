package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
)

type UpdateUserUseCase struct {
	userService services.UserService
}

func NewUpdateUserUseCase(us services.UserService) UpdateUserUseCase {
	return UpdateUserUseCase{userService: us}
}

func (uuuc *UpdateUserUseCase) Execute(dto dtos.UpdateSelfUserDTO) (*dtos.UserDTO, error) {
	usersFilterDTO := dtos.UsersFilterDTO{Ids: []int{dto.Id}}
	user, err := uuuc.userService.GetUser(usersFilterDTO)

	if err != nil {
		return nil, err
	}

	updateUserDTO := dtos.UpdateUserDTO{
		Id: dto.Id,
		//Organization: user.Organization,
		//Position:     user.Position,
		Role:      user.Role,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}
	updatedUser, err := uuuc.userService.UpdateUser(updateUserDTO)

	return updatedUser, err
}
