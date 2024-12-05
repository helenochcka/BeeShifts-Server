package ifaces

import (
	"BeeShifts-Server/internal/core/positions"
)

type PositionRepo interface {
	GetAll(filter positions.FilterDTO) ([]positions.Entity, error)
	GetOne(filter positions.FilterDTO) (*positions.Entity, error)
	Insert(position positions.Entity) (*positions.Entity, error)
	Update(position positions.Entity) (*positions.Entity, error)
}
