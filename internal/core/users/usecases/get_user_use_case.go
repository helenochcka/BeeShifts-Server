package usecases

import (
	"BeeShifts-Server/internal/core/organizations"
	orgsServices "BeeShifts-Server/internal/core/organizations/services"
	"BeeShifts-Server/internal/core/positions"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	"BeeShifts-Server/internal/core/users"
	usersServices "BeeShifts-Server/internal/core/users/services"
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
	user, err := guuc.userService.GetUser(filter)
	if err != nil {
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

	if user.OrganizationId != nil {
		orgFilter := organizations.FilterDTO{
			Ids: []int{*user.OrganizationId},
		}
		organization, err := guuc.orgService.GetOrganization(orgFilter)
		if err != nil {
			return nil, err
		}
		userDTO.Organization = &organization.Name
	}

	if user.PositionId != nil {
		positionFilter := positions.FilterDTO{
			Ids: []int{*user.PositionId},
		}
		position, err := guuc.positionService.GetPosition(positionFilter)
		if err != nil {
			return nil, err
		}
		userDTO.Position = &position.Name
	}

	return &userDTO, err
}
