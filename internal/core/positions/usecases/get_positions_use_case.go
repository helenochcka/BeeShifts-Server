package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
)

type GetPositionsUseCase struct {
	positionService services.PositionService
}

func NewGetPositionsUseCase(ps services.PositionService) GetPositionsUseCase {
	return GetPositionsUseCase{positionService: ps}
}

func (gpuc *GetPositionsUseCase) Execute(managerId int, filter positions.FilterDTO) ([]positions.Entity, error) {
	filter.ManagerIds = []int{managerId}

	positionEntities, err := gpuc.positionService.GetPositions(filter)
	if err != nil {
		return nil, err
	}
	return positionEntities, nil
}
