package usecases

import (
	"BeeShifts-Server/internal/core/organizations"
	"BeeShifts-Server/internal/core/organizations/services"
	"log/slog"
)

type GetOrgsUseCase struct {
	orgService services.OrgService
}

func NewGetOrgsUseCase(os services.OrgService) GetOrgsUseCase {
	return GetOrgsUseCase{orgService: os}
}

func (gouc *GetOrgsUseCase) Execute(filter organizations.FilterDTO) ([]organizations.Entity, error) {
	slog.Info("Getting organizations by filter...", "filter", filter)
	orgs, err := gouc.orgService.GetOrganizations(filter)

	if err != nil {
		slog.Error("Error getting organizations...", "error", err)
		return nil, err
	}

	return orgs, nil
}
