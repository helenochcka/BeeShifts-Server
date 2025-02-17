package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type AttachUserUseCase struct {
	userService     services.UserService
	positionService positionsServices.PositionService
}

func NewAttachUserUseCase(us services.UserService, ps positionsServices.PositionService) AttachUserUseCase {
	return AttachUserUseCase{userService: us, positionService: ps}
}

func (auuc *AttachUserUseCase) Execute(managerId int, dto users.AttachDTO) (*users.Entity, error) {
	slog.Info("Getting user to attach by id...", "id", dto.Id)
	employeeToAttach, err := auuc.userService.GetUser(users.FilterDTO{Ids: []int{dto.Id}})

	if err != nil {
		slog.Error("Error getting user to attach...", "err", err)
		return nil, err
	}

	slog.Info("Validating user's role...", "userRole", employeeToAttach.Role)
	if employeeToAttach.Role != users.Employee {
		slog.Error("Insufficient rights to attach user...", "user", employeeToAttach)
		return nil, users.InsufficientRights
	}

	slog.Info("Getting current manager by id...", "id", managerId)
	manager, err := auuc.userService.GetUser(users.FilterDTO{Ids: []int{managerId}})

	if err != nil {
		slog.Error("Error getting current manager...", "err", err)
		return nil, err
	}

	slog.Info("Validating employee's organization...", "user", employeeToAttach)
	if employeeToAttach.OrganizationId != nil && *employeeToAttach.OrganizationId != *manager.OrganizationId {
		slog.Error("Employee already attached to another organization...", "organization", *employeeToAttach.OrganizationId)
		return nil, users.EmployeeAlreadyAttached
	}

	positionFilter := positions.FilterDTO{
		Ids:        []int{dto.PositionId},
		ManagerIds: []int{managerId},
	}

	slog.Info("Checking if position belongs to current manager by filter...", "filter", positionFilter)
	_, err = auuc.positionService.GetPosition(positionFilter)

	if err != nil {
		slog.Error("Error getting position to attach...", "error", err)
		return nil, err
	}

	slog.Info("Changing organization and position of employee by dto...", "dto", dto)
	employeeToAttach.OrganizationId = manager.OrganizationId
	employeeToAttach.PositionId = &dto.PositionId

	slog.Info("Updating employee to attach...", "employeeToAttach", employeeToAttach)
	attachedUser, err := auuc.userService.UpdateUser(*employeeToAttach)

	if err != nil {
		slog.Error("Error updating employee to attach...", "err", err)
		return nil, err
	}

	return attachedUser, nil
}
