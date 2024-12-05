package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
)

type UpdateUserUseCase struct {
	userService services.UserService
}

func NewUpdateUserUseCase(us services.UserService) UpdateUserUseCase {
	return UpdateUserUseCase{userService: us}
}

func (uuuc *UpdateUserUseCase) Execute(id int, dto users.UpdateSelfDTO) (*users.Entity, error) {
	filter := users.FilterDTO{Ids: []int{id}}
	userToUpdate, err := uuuc.userService.GetUser(filter)

	if err != nil {
		return nil, err
	}

	filter = users.FilterDTO{Emails: []string{dto.Email}}
	conflictingUser, err := uuuc.userService.GetUser(filter)

	if conflictingUser != nil && conflictingUser.Id != userToUpdate.Id {
		return nil, users.EmailAlreadyUsed
	}

	userToUpdate.FirstName = dto.FirstName
	userToUpdate.LastName = dto.LastName
	userToUpdate.Email = dto.Email
	userToUpdate.Password = dto.Password

	updatedUser, err := uuuc.userService.UpdateUser(*userToUpdate)

	return updatedUser, err
}
