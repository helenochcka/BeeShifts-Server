package services

import (
	"BeeShifts-Server/internal/core/organizations"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"errors"
)

type OrgService struct {
	orgRepo ifaces.OrgRepo
}

func NewOrgService(or ifaces.OrgRepo) OrgService {
	return OrgService{orgRepo: or}
}

func (os *OrgService) GetOrganizations(filter organizations.FilterDTO) ([]organizations.Entity, error) {
	orgEntities, err := os.orgRepo.GetAll(filter)
	if err != nil {
		return nil, os.mapRepoErrToOrgsErr(err)
	}

	return orgEntities, nil
}

func (os *OrgService) GetOrganization(filter organizations.FilterDTO) (*organizations.Entity, error) {
	orgEntity, err := os.orgRepo.GetOne(filter)
	if err != nil {
		return nil, os.mapRepoErrToOrgsErr(err)
	}

	return orgEntity, nil
}

func (os *OrgService) mapRepoErrToOrgsErr(err error) error {
	switch {
	case errors.Is(err, repositories.MultipleRecFound):
		return organizations.MultipleOrgsFound
	case errors.Is(err, repositories.RecNotFound):
		return organizations.OrgNotFound
	default:
		return err
	}
}
