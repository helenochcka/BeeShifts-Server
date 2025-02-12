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

func (gpuc *GetPositionsUseCase) Execute(managerId int, dto positions.GetDTO) ([]positions.Entity, error) {

	filter := positions.FilterDTO{
		Ids:        dto.Ids,
		ManagerIds: []int{managerId},
		Names:      dto.Names,
	}

	slog.Info("Getting positions by dto...", "dto", dto)
	positionEntities, err := gpuc.positionService.GetPositions(filter)

	if err != nil {
		slog.Error("Error getting positions...", "error", err)
		return nil, err
	}

	return positionEntities, nil
}
