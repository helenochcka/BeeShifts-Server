package usecases

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type DetachUserUseCase struct {
	userService services.UserService
}

func NewDetachUserUseCase(us services.UserService) DetachUserUseCase {
	return DetachUserUseCase{userService: us}
}

func (duuc *DetachUserUseCase) Execute(managerId int, dto users.DetachDTO) (*users.Entity, error) {
	slog.Info("Getting user to detach by id...", "id", dto.Id)
	employeeToDetach, err := duuc.userService.GetUser(users.FilterDTO{Ids: []int{dto.Id}})

	if err != nil {
		slog.Error("Error getting user to detach...", "err", err)
		return nil, err
	}

	slog.Info("Validating user's role...", "userRole", employeeToDetach.Role)
	if employeeToDetach.Role != users.Employee {
		slog.Error("Insufficient rights to detach user...", "user", employeeToDetach)
		return nil, users.InsufficientRights
	}

	slog.Info("Getting current manager by id...", "id", managerId)
	manager, err := duuc.userService.GetUser(users.FilterDTO{Ids: []int{managerId}})

	if err != nil {
		slog.Error("Error getting current manager...", "err", err)
		return nil, err
	}

	slog.Info("Validating employee's organization...", "user", employeeToDetach)
	if employeeToDetach.OrganizationId == nil || *employeeToDetach.OrganizationId != *manager.OrganizationId {
		slog.Error("employee is not attached to manager's organization...",
			"employee's organization", *employeeToDetach.OrganizationId,
			"manager's organization", *manager.OrganizationId)
		return nil, users.EmployeeNotAttached
	}

	slog.Info("Setting employee's organization and position to null...")
	employeeToDetach.OrganizationId = nil
	employeeToDetach.PositionId = nil

	slog.Info("Updating employee to detach...", "employeeToDetach", employeeToDetach)
	detachedUser, err := duuc.userService.UpdateUser(*employeeToDetach)

	if err != nil {
		slog.Error("Error updating employee to detach...", "err", err)
		return nil, err
	}

	return detachedUser, nil
}
