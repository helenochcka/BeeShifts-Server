package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
	"log/slog"
)

type UpdatePositionUseCase struct {
	positionService services.PositionService
}

func NewUpdatePositionUseCase(ps services.PositionService) UpdatePositionUseCase {
	return UpdatePositionUseCase{positionService: ps}
}

func (upuc *UpdatePositionUseCase) Execute(managerId int, dto positions.UpdateDTO) (*positions.Entity, error) {

	filter := positions.FilterDTO{
		Ids:        []int{dto.Id},
		ManagerIds: []int{managerId},
	}

	slog.Info("Checking position by filter...", "filter", filter)
	positionToUpdate, err := upuc.positionService.GetPosition(filter)

	if err != nil {
		slog.Error("Error getting position to update...", "error", err)
		return nil, err
	}

	positionToUpdate.Name = dto.Name

	slog.Info("Updating position...", "positionToUpdate", positionToUpdate)
	updatedPosition, err := upuc.positionService.UpdatePosition(*positionToUpdate)

	if err != nil {
		slog.Error("Error updating position...", "error", err)
		return nil, err
	}

	return updatedPosition, nil
}
