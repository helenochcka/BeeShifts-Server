package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type UpdatePositionUseCase struct {
	positionService services.PositionService
}

func NewUpdatePositionUseCase(ps services.PositionService) UpdatePositionUseCase {
	return UpdatePositionUseCase{positionService: ps}
}

func (upuc *UpdatePositionUseCase) Execute(managerId int, dto dtos.UpdatePositionDTO) (*entities.PositionEntity, error) {
	positionToUpdate := entities.PositionEntity{
		Id:        dto.Id,
		ManagerId: managerId,
		Name:      dto.Name,
	}
	updatedPosition, err := upuc.positionService.UpdatePosition(positionToUpdate)

	return updatedPosition, err
}
