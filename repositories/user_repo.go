package repositories

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/repositories/models"
	"database/sql"
	_ "database/sql"
	"fmt"
	"strings"
)

type UserRepo interface {
	GetAll(filter dtos.UsersFilterDTO) ([]models.User, error)
	GetOne(filter dtos.UsersFilterDTO) (*models.User, error)
	Insert(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
}
type UserRepoPgSQL struct {
}

func NewUserRepoPgSQL() UserRepo {
	return &UserRepoPgSQL{}
}

func (ur *UserRepoPgSQL) GetAll(filter dtos.UsersFilterDTO) ([]models.User, error) {
	queryBase := "SELECT id, organization_id, position_id, role, first_name, last_name, email, password FROM users"

	conditions, args := ur.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var organizationId sql.NullInt64
		var positionId sql.NullInt64
		if err := rows.Scan(&user.Id, &organizationId, &positionId, &user.Role, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		if organizationId.Valid {
			value := int(organizationId.Int64)
			user.OrganizationId = &value
		} else {
			user.OrganizationId = nil
		}
		if positionId.Valid {
			value := int(positionId.Int64)
			user.PositionId = &value
		} else {
			user.PositionId = nil
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepoPgSQL) GetOne(filter dtos.UsersFilterDTO) (*models.User, error) {
	queryBase := "SELECT id, organization_id, position_id, role, first_name, last_name, email, password FROM users"

	conditions, args := ur.buildQueryParams(filter)

	query := queryBase
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	var organizationId sql.NullInt64
	var positionId sql.NullInt64
	if rows.Next() {
		if err := rows.Scan(&user.Id, &organizationId, &positionId, &user.Role, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			return nil, err
		}
	} else {
		return nil, RecNotFound
	}

	if rows.Next() {
		return nil, MultipleRecFound
	}

	if organizationId.Valid {
		value := int(organizationId.Int64)
		user.OrganizationId = &value
	} else {
		user.OrganizationId = nil
	}
	if positionId.Valid {
		value := int(positionId.Int64)
		user.PositionId = &value
	} else {
		user.PositionId = nil
	}

	return &user, nil
}

func (ur *UserRepoPgSQL) Insert(user models.User) (*models.User, error) {
	var userId int

	stmt := "insert into users (organization_id, position_id, role, first_name, last_name, email, password) values ($1, $2, $3, $4, $5, $6, $7) returning id"
	err := DB.QueryRow(stmt, user.OrganizationId, user.PositionId, user.Role, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return nil, err
	}

	insertedUser, err := ur.GetOne(dtos.UsersFilterDTO{Ids: []int{userId}})
	if err != nil {
		return nil, err
	}

	return insertedUser, nil
}

func (ur *UserRepoPgSQL) Update(user models.User) (*models.User, error) {

	stmt := "update users set organization_id=$1, position_id=$2, role=$3, first_name=$4, last_name=$5,  email=$6, password=$7 where id = $8"
	_, err := DB.Exec(stmt, user.OrganizationId, user.PositionId, user.Role, user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := ur.GetOne(dtos.UsersFilterDTO{Ids: []int{user.Id}})
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (ur *UserRepoPgSQL) buildQueryParams(filter dtos.UsersFilterDTO) ([]string, []interface{}) {
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
