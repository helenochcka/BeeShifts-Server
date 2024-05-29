package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type UpdateUserUseCase struct {
	userService services.UserService
}

func NewUpdateUserUseCase(us services.UserService) UpdateUserUseCase {
	return UpdateUserUseCase{userService: us}
}

func (uuuc *UpdateUserUseCase) Execute(id int, dto dtos.UpdateSelfUserDTO) (*entities.UserEntity, error) {
	usersFilterDTO := dtos.UsersFilterDTO{Ids: []int{id}}
	user, err := uuuc.userService.GetUser(usersFilterDTO)

	if err != nil {
		return nil, err
	}

	userToUpdate := entities.UserEntity{
		Id:             id,
		OrganizationId: user.OrganizationId,
		PositionId:     user.PositionId,
		Role:           user.Role,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
	}
	updatedUser, err := uuuc.userService.UpdateUser(userToUpdate)

	return updatedUser, err
}
