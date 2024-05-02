package repositories

import (
	"BeeShifts-Server/models"
	_ "database/sql"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur *UserRepository) Add(user models.User) (models.User, error) {
	var userId int

	stmt := "insert into users (first_name, last_name, organization_id, position_id, email, password) values ($1, $2, $3, $4, $5, $6) returning id"
	err := DB.QueryRow(stmt, user.FirstName, user.LastName, user.Organization, user.Position, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return models.User{}, err
	}

	user.Id = userId

	return user, nil
}

func (ur *UserRepository) GetAll() ([]models.User, error) {

	stmt := "select id, first_name, last_name, organization_id, position_id, email, password from users"
	rows, err := DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Organization, &user.Position, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, err
}

func (ur *UserRepository) GetByID(id int) (models.User, error) {

	stmt := "select id, first_name, last_name, organization_id, position_id, email, password from users where id = $1"
	row := DB.QueryRow(stmt, id)

	user := models.User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Organization, &user.Position, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (models.User, error) {

	stmt := "select id, first_name, last_name, organization_id, position_id, email, password from users where email = $1"
	row := DB.QueryRow(stmt, email)

	user := models.User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Organization, &user.Position, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Update(user models.User) (models.User, error) {

	stmt := "update users set first_name=$1, last_name=$2, organization_id=$3, position_id=$4, email=$5, password=$6 where id = $7"
	_, err := DB.Exec(stmt, user.FirstName, user.LastName, user.Organization, user.Position, user.Email, user.Password, user.Id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(id int) (models.User, error) {

	user, _ := ur.GetByID(id)

	stmt := "delete from users where id = $1"
	_, err := DB.Exec(stmt, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
