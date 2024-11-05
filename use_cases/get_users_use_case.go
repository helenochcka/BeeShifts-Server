package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
)

type GetUsersUseCase struct {
	userService     services.UserService
	orgService      services.OrgService
	positionService services.PositionService
}

func NewGetUsersUseCase(us services.UserService, os services.OrgService, ps services.PositionService) GetUsersUseCase {
	return GetUsersUseCase{userService: us, orgService: os, positionService: ps}
}

func (guuc *GetUsersUseCase) Execute(dto dtos.UsersFilterDTO) ([]dtos.UserDTO, error) {
	users, err := guuc.userService.GetUsers(dto)
	if err != nil {
		return nil, err
	}

	var organizationFilter dtos.OrgsFilterDTO
	var positionFilter dtos.PositionsFilterDTO

	for _, user := range users {
		if user.OrganizationId != nil {
			organizationFilter.Ids = append(organizationFilter.Ids, *user.OrganizationId)
		}

		if user.PositionId != nil {
			positionFilter.Ids = append(positionFilter.Ids, *user.PositionId)
		}
	}

	orgEntities, err := guuc.orgService.GetOrganizations(organizationFilter)
	if err != nil {
		return nil, err
	}

	orgIdToOrgNameMap := make(map[int]string)
	for _, orgEntity := range orgEntities {
		orgIdToOrgNameMap[orgEntity.Id] = orgEntity.Name
	}

	positionEntities, err := guuc.positionService.GetPositions(positionFilter)
	if err != nil {
		return nil, err
	}

	positionIdToPositionNameMap := make(map[int]string)
	for _, positionEntity := range positionEntities {
		positionIdToPositionNameMap[positionEntity.Id] = positionEntity.Name
	}

	var userDTOS []dtos.UserDTO
	for _, user := range users {
		userDTO := dtos.UserDTO{
			Id:           user.Id,
			Organization: nil,
			Position:     nil,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
		}
		if user.OrganizationId != nil {
			orgName := orgIdToOrgNameMap[*(user.OrganizationId)]
			userDTO.Organization = &orgName
		}

		if user.PositionId != nil {
			positionName := positionIdToPositionNameMap[*(user.PositionId)]
			userDTO.Position = &positionName
		}

		userDTOS = append(userDTOS, userDTO)
	}

	return userDTOS, err
}
