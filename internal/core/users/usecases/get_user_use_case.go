package usecases

import (
	"BeeShifts-Server/internal/core/organizations"
	orgsServices "BeeShifts-Server/internal/core/organizations/services"
	"BeeShifts-Server/internal/core/positions"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	"BeeShifts-Server/internal/core/users"
	usersServices "BeeShifts-Server/internal/core/users/services"
	"log/slog"
)

type GetUserUseCase struct {
	userService     usersServices.UserService
	orgService      orgsServices.OrgService
	positionService positionsServices.PositionService
}

func NewGetUserUseCase(us usersServices.UserService, os orgsServices.OrgService, ps positionsServices.PositionService) GetUserUseCase {
	return GetUserUseCase{userService: us, orgService: os, positionService: ps}
}

func (guuc *GetUserUseCase) Execute(filter users.FilterDTO) (*users.ViewDTO, error) {
	slog.Info("Getting user by filter...", "filter", filter)
	user, err := guuc.userService.GetUser(filter)

	if err != nil {
		slog.Error("Error getting user...", "err", err)
		return nil, err
	}

	userDTO := users.ViewDTO{
		Id:           user.Id,
		Organization: nil,
		Position:     nil,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
	}

	slog.Info("Checking if organization id present in user...", "user", user)
	if user.OrganizationId != nil {
		orgFilter := organizations.FilterDTO{
			Ids: []int{*user.OrganizationId},
		}

		slog.Info("Getting organization by filter...", "filter", orgFilter)
		organization, err := guuc.orgService.GetOrganization(orgFilter)

		if err != nil {
			slog.Error("Error getting organization...", "err", err)
			return nil, err
		}

		userDTO.Organization = &organization.Name
	}

	slog.Info("Checking if position id present in user...", "user", user)
	if user.PositionId != nil {
		positionFilter := positions.FilterDTO{
			Ids: []int{*user.PositionId},
		}

		slog.Info("Getting position by filter...", "filter", positionFilter)
		position, err := guuc.positionService.GetPosition(positionFilter)

		if err != nil {
			slog.Error("Error getting position...", "err", err)
			return nil, err
		}

		userDTO.Position = &position.Name
	}

	return &userDTO, nil
}
