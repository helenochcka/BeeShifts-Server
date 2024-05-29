package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/services"
)

type CreatePositionUseCase struct {
	positionService services.PositionService
}

func NewCreatePositionUseCase(ps services.PositionService) CreatePositionUseCase {
	return CreatePositionUseCase{positionService: ps}
}

func (cpuc *CreatePositionUseCase) Execute(managerId int, dto dtos.CreatePositionDTO) (*entities.PositionEntity, error) {
	positionToCreate := entities.PositionEntity{
		ManagerId: managerId,
		Name:      dto.Name,
	}
	user, err := cpuc.positionService.CreatePosition(positionToCreate)

	return user, err
}
