package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type GetPositionsUseCase struct {
	positionService services.PositionService
}

func NewGetPositionsUseCase(ps services.PositionService) GetPositionsUseCase {
	return GetPositionsUseCase{positionService: ps}
}

func (gpuc *GetPositionsUseCase) Execute(dto dtos.PositionsFilterDTO) ([]entities.PositionEntity, error) {
	positions, err := gpuc.positionService.GetPositions(dto)
	if err != nil {
		return nil, err
	}
	return positions, nil
}
