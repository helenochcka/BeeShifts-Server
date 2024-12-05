package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
)

type CreatePositionUseCase struct {
	positionService services.PositionService
}

func NewCreatePositionUseCase(ps services.PositionService) CreatePositionUseCase {
	return CreatePositionUseCase{positionService: ps}
}

func (cpuc *CreatePositionUseCase) Execute(managerId int, dto positions.CreateDTO) (*positions.Entity, error) {
	positionToCreate := positions.Entity{
		ManagerId: managerId,
		Name:      dto.Name,
	}
	createdPosition, err := cpuc.positionService.CreatePosition(positionToCreate)

	return createdPosition, err
}
