package postgres

import (
	"BeeShifts-Server/internal/core/organizations"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"BeeShifts-Server/pkg/db"
	"fmt"
	"strings"
)

type OrgRepoPgSQL struct {
}

func NewOrgRepoPgSQL() ifaces.OrgRepo {
	return &OrgRepoPgSQL{}
}

func (or *OrgRepoPgSQL) GetAll(filter organizations.FilterDTO) ([]organizations.Entity, error) {
	queryBase := "SELECT id, name FROM organizations"

	conditions, args := or.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entities []organizations.Entity
	for rows.Next() {
		var entity organizations.Entity
		if err = rows.Scan(&entity.Id, &entity.Name); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (or *OrgRepoPgSQL) GetOne(filter organizations.FilterDTO) (*organizations.Entity, error) {
	queryBase := "SELECT id, name FROM organizations"

	conditions, args := or.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entity organizations.Entity
	if rows.Next() {
		if err = rows.Scan(&entity.Id, &entity.Name); err != nil {
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

func (or *OrgRepoPgSQL) buildQueryParams(filter organizations.FilterDTO) ([]string, []interface{}) {
	var conditions []string
	var args []interface{}

	if len(filter.Ids) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "id", placeholders(len(filter.Ids), len(args)+1)))
		for _, arg := range filter.Ids {
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
