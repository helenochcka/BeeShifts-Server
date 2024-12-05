package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
)

type UpdatePositionUseCase struct {
	positionService services.PositionService
}

func NewUpdatePositionUseCase(ps services.PositionService) UpdatePositionUseCase {
	return UpdatePositionUseCase{positionService: ps}
}

func (upuc *UpdatePositionUseCase) Execute(managerId int, dto positions.UpdateDTO) (*positions.Entity, error) {
	positionToUpdate := positions.Entity{
		Id:        dto.Id,
		ManagerId: managerId,
		Name:      dto.Name,
	}
	updatedPosition, err := upuc.positionService.UpdatePosition(positionToUpdate)

	return updatedPosition, err
}
