package usecases

import (
	"BeeShifts-Server/internal/core/organizations"
	"BeeShifts-Server/internal/core/organizations/services"
)

type GetOrgsUseCase struct {
	orgService services.OrgService
}

func NewGetOrgsUseCase(os services.OrgService) GetOrgsUseCase {
	return GetOrgsUseCase{orgService: os}
}

func (gouc *GetOrgsUseCase) Execute(filter organizations.FilterDTO) ([]organizations.Entity, error) {
	orgs, err := gouc.orgService.GetOrganizations(filter)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
