package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
	"log/slog"
)

type GetPositionsUseCase struct {
	positionService services.PositionService
}

func NewGetPositionsUseCase(ps services.PositionService) GetPositionsUseCase {
	return GetPositionsUseCase{positionService: ps}
}

func (gpuc *GetPositionsUseCase) Execute(managerId int, filter positions.FilterDTO) ([]positions.Entity, error) {
	filter.ManagerIds = []int{managerId}

	slog.Info("Getting positions by filter...", "filter", filter)
	positionEntities, err := gpuc.positionService.GetPositions(filter)

	if err != nil {
		slog.Error("Error getting positions...", "error", err)
		return nil, err
	}

	return positionEntities, nil
}
