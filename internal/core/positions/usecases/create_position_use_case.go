package usecases

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/services"
	"log/slog"
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
	slog.Info("Creating position...", "positionToCreate", positionToCreate)
	createdPosition, err := cpuc.positionService.CreatePosition(positionToCreate)

	if err != nil {
		slog.Error("Error creating position...", "error", err)
		return nil, err
	}

	return createdPosition, nil
}
