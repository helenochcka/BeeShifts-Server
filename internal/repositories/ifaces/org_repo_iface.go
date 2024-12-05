package ifaces

import "BeeShifts-Server/internal/core/organizations"

type OrgRepo interface {
	GetAll(filter organizations.FilterDTO) ([]organizations.Entity, error)
	GetOne(filter organizations.FilterDTO) (*organizations.Entity, error)
}
