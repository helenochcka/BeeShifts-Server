package repositories

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/repositories/models"
	"fmt"
	"strings"
)

type PositionRepo struct {
}

func NewPositionRepo() PositionRepo {
	return PositionRepo{}
}

func (pr *PositionRepo) GetAll(filter dtos.GetPositionsDTO) ([]models.Position, error) {
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

	var positions []models.Position
	for rows.Next() {
		var position models.Position
		if err := rows.Scan(&position.Id, &position.ManagerID, &position.Name); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}

func (pr *PositionRepo) GetOne(filter dtos.GetPositionsDTO) (*models.Position, error) {
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

	var position models.Position
	if rows.Next() {
		if err := rows.Scan(&position.Id, &position.ManagerID, &position.Name); err != nil {
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

func (pr *PositionRepo) buildQueryParams(filter dtos.GetPositionsDTO) ([]string, []interface{}) {
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
