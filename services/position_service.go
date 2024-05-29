package services

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/repositories"
)

type PositionService struct {
	positionRepo repositories.PositionRepo
}

func NewPositionService(pr repositories.PositionRepo) PositionService {
	return PositionService{positionRepo: pr}
}

func (ps *PositionService) GetPositions(dto dtos.PositionsFilterDTO) ([]entities.PositionEntity, error) {
	positionEntities, err := ps.positionRepo.GetAll(dto)

	return positionEntities, err
}

func (ps *PositionService) GetPosition(dto dtos.PositionsFilterDTO) (*entities.PositionEntity, error) {
	positionEntity, err := ps.positionRepo.GetOne(dto)
	if err != nil {
		return nil, err
	} //TODO cast repo err to domain err

	return positionEntity, nil
}

func (ps *PositionService) CreatePosition(positionToCreate entities.PositionEntity) (*entities.PositionEntity, error) {
	createdPosition, err := ps.positionRepo.Insert(positionToCreate)
	if err != nil {
		return nil, err
	}
	return createdPosition, nil
}

func (ps *PositionService) UpdatePosition(positionToUpdate entities.PositionEntity) (*entities.PositionEntity, error) {
	updatedPosition, err := ps.positionRepo.Update(positionToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedPosition, nil
}
