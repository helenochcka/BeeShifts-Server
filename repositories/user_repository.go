package repositories

import (
	"BeeShifts-Server/models"
	_ "database/sql"
	"strings"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

type UserFilter struct {
	Ids             []interface{}
	OrganizationIds []interface{}
	PositionIds     []interface{}
	RoleIds         []interface{}
	FirstNames      []interface{}
	LastNames       []interface{}
	Emails          []interface{}
	Passwords       []interface{}
}

func (ur *UserRepository) Get(filter UserFilter) ([]models.User, error) {
	queryBase := "SELECT id, organization_id, position_id, role_id, first_name, last_name, email, password FROM users"

	var conditions []string
	var args []interface{}
	buildFilter(filter.Ids, "id", &conditions, &args)
	buildFilter(filter.OrganizationIds, "organization_id", &conditions, &args)
	buildFilter(filter.PositionIds, "position_id", &conditions, &args)
	buildFilter(filter.RoleIds, "role_id", &conditions, &args)
	buildFilter(filter.FirstNames, "first_name", &conditions, &args)
	buildFilter(filter.LastNames, "last_name", &conditions, &args)
	buildFilter(filter.Emails, "email", &conditions, &args)
	buildFilter(filter.Passwords, "password", &conditions, &args)

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
		if err := rows.Scan(&user.Id, &user.Organization, &user.Position, &user.Role, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) Add(user models.User) (*models.User, error) {
	var userId int

	stmt := "insert into users (organization_id, position_id, role_id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5, $6, $7) returning id"
	err := DB.QueryRow(stmt, user.Organization, user.Position, user.Role, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return nil, err
	}

	user.Id = userId

	return &user, nil
}

func (ur *UserRepository) Update(user models.User) (*models.User, error) {

	stmt := "update users set organization_id=$1, position_id=$2, role_id=$3, first_name=$4, last_name=$5,  email=$6, password=$7 where id = $8"
	_, err := DB.Exec(stmt, user.Organization, user.Position, user.Role, user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Delete(id int) error {

	stmt := "delete from users where id = $1"
	_, err := DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
