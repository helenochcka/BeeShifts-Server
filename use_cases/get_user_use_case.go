package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
)

type GetUserUseCase struct {
	userService     services.UserService
	orgService      services.OrgService
	positionService services.PositionService
}

func NewGetUserUseCase(us services.UserService, os services.OrgService, ps services.PositionService) GetUserUseCase {
	return GetUserUseCase{userService: us, orgService: os, positionService: ps}
}

func (guuc *GetUserUseCase) Execute(dto dtos.UsersFilterDTO) (*dtos.UserDTO, error) {
	user, err := guuc.userService.GetUser(dto)
	if err != nil {
		return nil, err
	}

	userDTO := dtos.UserDTO{
		Id:           user.Id,
		Organization: nil,
		Position:     nil,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
	}

	if user.OrganizationId != nil {
		organizationFilter := dtos.OrgsFilterDTO{
			Ids: []int{*user.OrganizationId},
		}
		organization, err := guuc.orgService.GetOrganization(organizationFilter)
		if err != nil {
			return nil, err
		}
		userDTO.Organization = &organization.Name
	}

	if user.PositionId != nil {
		positionFilter := dtos.PositionsFilterDTO{
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
