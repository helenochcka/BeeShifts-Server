package postgres

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"BeeShifts-Server/pkg/db"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type UserRepoPgSQL struct {
}

func NewUserRepoPgSQL() ifaces.UserRepo {
	return &UserRepoPgSQL{}
}

func (ur *UserRepoPgSQL) GetAll(filter users.FilterDTO) ([]users.Entity, error) {
	queryBase := "SELECT id, organization_id, position_id, role, first_name, last_name, email, password FROM users"

	conditions, args := ur.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entities []users.Entity
	for rows.Next() {
		var entity users.Entity
		var orgId sql.NullInt64
		var positionId sql.NullInt64
		if err = rows.Scan(&entity.Id, &orgId, &positionId, &entity.Role, &entity.FirstName, &entity.LastName, &entity.Email, &entity.Password); err != nil {
			return nil, err
		}
		if orgId.Valid {
			value := int(orgId.Int64)
			entity.OrganizationId = &value
		} else {
			entity.OrganizationId = nil
		}
		if positionId.Valid {
			value := int(positionId.Int64)
			entity.PositionId = &value
		} else {
			entity.PositionId = nil
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (ur *UserRepoPgSQL) GetOne(filter users.FilterDTO) (*users.Entity, error) {
	queryBase := "SELECT id, organization_id, position_id, role, first_name, last_name, email, password FROM users"

	conditions, args := ur.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var entity users.Entity
	var orgId sql.NullInt64
	var positionId sql.NullInt64
	if rows.Next() {
		if err = rows.Scan(&entity.Id, &orgId, &positionId, &entity.Role, &entity.FirstName, &entity.LastName, &entity.Email, &entity.Password); err != nil {
			return nil, err
		}
	} else {
		return nil, repositories.RecNotFound
	}

	if rows.Next() {
		return nil, repositories.MultipleRecFound
	}

	if orgId.Valid {
		value := int(orgId.Int64)
		entity.OrganizationId = &value
	} else {
		entity.OrganizationId = nil
	}
	if positionId.Valid {
		value := int(positionId.Int64)
		entity.PositionId = &value
	} else {
		entity.PositionId = nil
	}

	return &entity, nil
}

func (ur *UserRepoPgSQL) GetOneOrNil(filter users.FilterDTO) (*users.Entity, error) {
	entity, err := ur.GetOne(filter)

	if errors.Is(err, users.UserNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (ur *UserRepoPgSQL) Insert(entity users.Entity) (*users.Entity, error) {
	var id int

	stmt := "insert into users (organization_id, position_id, role, first_name, last_name, email, password) values ($1, $2, $3, $4, $5, $6, $7) returning id"
	err := db.DB.QueryRow(stmt, entity.OrganizationId, entity.PositionId, entity.Role, entity.FirstName, entity.LastName, entity.Email, entity.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	insertedEntity, err := ur.GetOne(users.FilterDTO{Ids: []int{id}})
	if err != nil {
		return nil, err
	}

	return insertedEntity, nil
}

func (ur *UserRepoPgSQL) Update(entity users.Entity) (*users.Entity, error) {

	stmt := "update users set organization_id=$1, position_id=$2, role=$3, first_name=$4, last_name=$5,  email=$6, password=$7 where id = $8"
	_, err := db.DB.Exec(stmt, entity.OrganizationId, entity.PositionId, entity.Role, entity.FirstName, entity.LastName, entity.Email, entity.Password, entity.Id)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := ur.GetOne(users.FilterDTO{Ids: []int{entity.Id}})
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (ur *UserRepoPgSQL) buildQueryParams(filter users.FilterDTO) ([]string, []interface{}) {
	var conditions []string
	var args []interface{}

	if len(filter.Ids) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "id", placeholders(len(filter.Ids), len(args)+1)))
		for _, arg := range filter.Ids {
			args = append(args, arg)
		}
	}

	if len(filter.OrganizationIds) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "organization_id", placeholders(len(filter.OrganizationIds), len(args)+1)))
		for _, arg := range filter.OrganizationIds {
			args = append(args, arg)
		}
	}

	if len(filter.PositionIds) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "position_id", placeholders(len(filter.PositionIds), len(args)+1)))
		for _, arg := range filter.PositionIds {
			args = append(args, arg)
		}
	}

	if len(filter.FirstNames) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "first_name", placeholders(len(filter.FirstNames), len(args)+1)))
		for _, arg := range filter.FirstNames {
			args = append(args, arg)
		}
	}

	if len(filter.LastNames) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "last_name", placeholders(len(filter.LastNames), len(args)+1)))
		for _, arg := range filter.LastNames {
			args = append(args, arg)
		}
	}

	if len(filter.Emails) > 0 {
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", "email", placeholders(len(filter.Emails), len(args)+1)))
		for _, arg := range filter.Emails {
			args = append(args, arg)
		}
	}

	return conditions, args
}
