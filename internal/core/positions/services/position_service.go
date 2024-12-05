package services

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"errors"
)

type PositionService struct {
	positionRepo ifaces.PositionRepo
}

func NewPositionService(pr ifaces.PositionRepo) PositionService {
	return PositionService{positionRepo: pr}
}

func (ps *PositionService) GetPositions(filter positions.FilterDTO) ([]positions.Entity, error) {
	positionEntities, err := ps.positionRepo.GetAll(filter)
	if err != nil {
		return nil, ps.mapRepoErrToPositionsErr(err)
	}

	return positionEntities, nil
}

func (ps *PositionService) GetPosition(filter positions.FilterDTO) (*positions.Entity, error) {
	positionEntity, err := ps.positionRepo.GetOne(filter)
	if err != nil {
		return nil, ps.mapRepoErrToPositionsErr(err)
	}

	return positionEntity, nil
}

func (ps *PositionService) CreatePosition(positionToCreate positions.Entity) (*positions.Entity, error) {
	createdPosition, err := ps.positionRepo.Insert(positionToCreate)
	if err != nil {
		return nil, ps.mapRepoErrToPositionsErr(err)
	}
	return createdPosition, nil
}

func (ps *PositionService) UpdatePosition(positionToUpdate positions.Entity) (*positions.Entity, error) {
	updatedPosition, err := ps.positionRepo.Update(positionToUpdate)
	if err != nil {
		return nil, ps.mapRepoErrToPositionsErr(err)
	}

	return updatedPosition, nil
}

func (ps *PositionService) mapRepoErrToPositionsErr(err error) error {
	switch {
	case errors.Is(err, repositories.MultipleRecFound):
		return positions.MultiplePositionsFound
	case errors.Is(err, repositories.RecNotFound):
		return positions.PositionNotFound
	default:
		return err
	}
}
