package services

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/repositories"
)

type OrgService struct {
	orgRepo repositories.OrgRepo
}

func NewOrgService(or repositories.OrgRepo) OrgService {
	return OrgService{orgRepo: or}
}

func (os *OrgService) GetOrganizations(dto dtos.OrgsFilterDTO) ([]entities.OrganizationEntity, error) {
	orgEntities, err := os.orgRepo.GetAll(dto)

	return orgEntities, err
}

func (os *OrgService) GetOrganization(dto dtos.OrgsFilterDTO) (*entities.OrganizationEntity, error) {
	orgEntity, err := os.orgRepo.GetOne(dto)
	if err != nil {
		return nil, err
	} //TODO cast repo err to domain err

	return orgEntity, nil
}
