package ifaces

import (
	"BeeShifts-Server/internal/core/users"
)

type UserRepo interface {
	GetAll(filter users.FilterDTO) ([]users.Entity, error)
	GetOne(filter users.FilterDTO) (*users.Entity, error)
	Insert(user users.Entity) (*users.Entity, error)
	Update(user users.Entity) (*users.Entity, error)
}
