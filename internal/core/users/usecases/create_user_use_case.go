package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
)

type CreateUserUseCase struct {
	userService services.UserService
}

func NewCreateUserUseCase(us services.UserService) CreateUserUseCase {
	return CreateUserUseCase{userService: us}
}

func (cuuc *CreateUserUseCase) Execute(dto users.CreateDTO) (*users.Entity, error) {
	filter := users.FilterDTO{
		Emails: []string{dto.Email},
	}
	conflictingUser, err := cuuc.userService.GetUser(filter)
	if conflictingUser != nil {
		return nil, users.EmailAlreadyUsed
	}

	if !cuuc.userService.IsRoleValid(dto.Role) {
		return nil, users.RoleDoesNotExist
	}

	userToCreate := users.Entity{
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
