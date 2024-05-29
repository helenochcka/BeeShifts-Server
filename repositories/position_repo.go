package repositories

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"fmt"
	"strings"
)

type PositionRepo interface {
	GetAll(filter dtos.PositionsFilterDTO) ([]entities.PositionEntity, error)
	GetOne(filter dtos.PositionsFilterDTO) (*entities.PositionEntity, error)
	Insert(position entities.PositionEntity) (*entities.PositionEntity, error)
	Update(position entities.PositionEntity) (*entities.PositionEntity, error)
}

type PositionRepoPgSQL struct {
}

func NewPositionRepoPgSQL() PositionRepo {
	return &PositionRepoPgSQL{}
}

func (pr *PositionRepoPgSQL) GetAll(filter dtos.PositionsFilterDTO) ([]entities.PositionEntity, error) {
	queryBase := "SELECT id, manager_id, name FROM positions"

	conditions, args := pr.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []entities.PositionEntity
	for rows.Next() {
		var position entities.PositionEntity
		if err := rows.Scan(&position.Id, &position.ManagerId, &position.Name); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}

func (pr *PositionRepoPgSQL) GetOne(filter dtos.PositionsFilterDTO) (*entities.PositionEntity, error) {
	queryBase := "SELECT id, manager_id, name FROM positions"

	conditions, args := pr.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var position entities.PositionEntity
	if rows.Next() {
		if err := rows.Scan(&position.Id, &position.ManagerId, &position.Name); err != nil {
			return nil, err
		}
	} else {
		return nil, RecNotFound
	}

	if rows.Next() {
		return nil, MultipleRecFound
	}

	return &position, nil
}

func (pr *PositionRepoPgSQL) Insert(position entities.PositionEntity) (*entities.PositionEntity, error) {
	var positionId int

	stmt := "insert into positions (name, manager_id) values ($1, $2) returning id"
	err := DB.QueryRow(stmt, position.Name, position.ManagerId).Scan(&positionId)
	if err != nil {
		return nil, err
	}

	insertedPosition, err := pr.GetOne(dtos.PositionsFilterDTO{Ids: []int{positionId}})
	if err != nil {
		return nil, err
	}

	return insertedPosition, nil
}

func (pr *PositionRepoPgSQL) Update(position entities.PositionEntity) (*entities.PositionEntity, error) {

	stmt := "update positions set name=$1, manager_id=$2 where id = $3"
	_, err := DB.Exec(stmt, position.Name, position.ManagerId, position.Id)
	if err != nil {
		return nil, err
	}

	updatedPosition, err := pr.GetOne(dtos.PositionsFilterDTO{Ids: []int{position.Id}})
	if err != nil {
		return nil, err
	}

	return updatedPosition, nil
}

func (pr *PositionRepoPgSQL) buildQueryParams(filter dtos.PositionsFilterDTO) ([]string, []interface{}) {
	var conditions []string
	var args []interface{}

	if len(filter.Ids) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "id", placeholders(len(filter.Ids), len(args)+1)))
		for _, arg := range filter.Ids {
			args = append(args, arg)
		}
	}

	if len(filter.ManagerIds) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "manager_id", placeholders(len(filter.ManagerIds), len(args)+1)))
		for _, arg := range filter.ManagerIds {
			args = append(args, arg)
		}
	}

	if len(filter.Names) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "name", placeholders(len(filter.Names), len(args)+1)))
		for _, arg := range filter.Names {
			args = append(args, arg)
		}
	}

	return conditions, args
}
