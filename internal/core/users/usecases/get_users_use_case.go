package usecases

import (
	"BeeShifts-Server/internal/core/organizations"
	orgsServices "BeeShifts-Server/internal/core/organizations/services"
	"BeeShifts-Server/internal/core/positions"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	"BeeShifts-Server/internal/core/users"
	usersServices "BeeShifts-Server/internal/core/users/services"
)

type GetUsersUseCase struct {
	userService     usersServices.UserService
	orgService      orgsServices.OrgService
	positionService positionsServices.PositionService
}

func NewGetUsersUseCase(us usersServices.UserService, os orgsServices.OrgService, ps positionsServices.PositionService) GetUsersUseCase {
	return GetUsersUseCase{userService: us, orgService: os, positionService: ps}
}

func (guuc *GetUsersUseCase) Execute(filter users.FilterDTO) ([]users.ViewDTO, error) {
	userEntities, err := guuc.userService.GetUsers(filter)
	if err != nil {
		return nil, err
	}

	var orgFilter organizations.FilterDTO
	var positionFilter positions.FilterDTO

	for _, userEntity := range userEntities {
		if userEntity.OrganizationId != nil {
			orgFilter.Ids = append(orgFilter.Ids, *userEntity.OrganizationId)
		}

		if userEntity.PositionId != nil {
			positionFilter.Ids = append(positionFilter.Ids, *userEntity.PositionId)
		}
	}

	orgEntities, err := guuc.orgService.GetOrganizations(orgFilter)
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

	var userDTOS []users.ViewDTO
	for _, user := range userEntities {
		userDTO := users.ViewDTO{
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
