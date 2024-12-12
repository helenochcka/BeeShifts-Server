package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
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

	slog.Info("Looking for conflicting user by filter...", "filter", filter)
	conflictingUser, err := cuuc.userService.FindUser(filter)

	if err != nil {
		slog.Error("Error getting conflicting user...", "error", err)
		return nil, err
	}

	if conflictingUser != nil {
		slog.Error("Conflicting by email user exists...", "email", dto.Email)
		return nil, users.EmailAlreadyUsed
	}

	slog.Info("Validating given role...", "role", dto.Role)
	if !cuuc.userService.IsRoleValid(dto.Role) {
		slog.Error("Given user role is invalid....", "role", dto.Role)
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

	slog.Info("Creating user by dto...", "dto", dto)
	createdUser, err := cuuc.userService.CreateUser(userToCreate)

	if err != nil {
		slog.Error("Error creating user...", "error", err)
		return nil, err
	}

	return createdUser, nil
}
