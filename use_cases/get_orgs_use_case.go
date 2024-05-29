package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type GetOrgsUseCase struct {
	orgService services.OrgService
}

func NewGetOrgsUseCase(os services.OrgService) GetOrgsUseCase {
	return GetOrgsUseCase{orgService: os}
}

func (gouc *GetOrgsUseCase) Execute(dto dtos.OrgsFilterDTO) ([]entities.OrganizationEntity, error) {
	orgs, err := gouc.orgService.GetOrganizations(dto)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
