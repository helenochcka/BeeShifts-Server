package postgres

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"BeeShifts-Server/pkg/db"
	"fmt"
	"strings"
)

type PositionRepoPgSQL struct {
}

func NewPositionRepoPgSQL() ifaces.PositionRepo {
	return &PositionRepoPgSQL{}
}

func (pr *PositionRepoPgSQL) GetAll(filter positions.FilterDTO) ([]positions.Entity, error) {
	queryBase := "SELECT id, manager_id, name FROM positions"

	conditions, args := pr.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entities []positions.Entity
	for rows.Next() {
		var entity positions.Entity
		if err = rows.Scan(&entity.Id, &entity.ManagerId, &entity.Name); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (pr *PositionRepoPgSQL) GetOne(filter positions.FilterDTO) (*positions.Entity, error) {
	queryBase := "SELECT id, manager_id, name FROM positions"

	conditions, args := pr.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entity positions.Entity
	if rows.Next() {
		if err = rows.Scan(&entity.Id, &entity.ManagerId, &entity.Name); err != nil {
			return nil, err
		}
	} else {
		return nil, repositories.RecNotFound
	}

	if rows.Next() {
		return nil, repositories.MultipleRecFound
	}

	return &entity, nil
}

func (pr *PositionRepoPgSQL) Insert(entity positions.Entity) (*positions.Entity, error) {
	var id int

	stmt := "insert into positions (name, manager_id) values ($1, $2) returning id"
	err := db.DB.QueryRow(stmt, entity.Name, entity.ManagerId).Scan(&id)
	if err != nil {
		return nil, err
	}

	insertedEntity, err := pr.GetOne(positions.FilterDTO{Ids: []int{id}})
	if err != nil {
		return nil, err
	}

	return insertedEntity, nil
}

func (pr *PositionRepoPgSQL) Update(entity positions.Entity) (*positions.Entity, error) {

	stmt := "update positions set name=$1, manager_id=$2 where id = $3"
	_, err := db.DB.Exec(stmt, entity.Name, entity.ManagerId, entity.Id)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := pr.GetOne(positions.FilterDTO{Ids: []int{entity.Id}})
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (pr *PositionRepoPgSQL) buildQueryParams(filter positions.FilterDTO) ([]string, []interface{}) {
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
