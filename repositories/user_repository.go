package repositories

import (
	"BeeShifts-Server/models"
	"database/sql"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur *UserRepository) Add(user models.User) (models.User, error) {
	tx, err := DB.Begin()
	if err != nil {
		return models.User{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO users (first_name, last_name, organization, position, email, password) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.FirstName, user.LastName, user.Organization, user.Position, user.Email, user.Password)

	userId, _ := res.LastInsertId()
	user.Id = int(userId)

	if err != nil {
		return models.User{}, err
	}

	tx.Commit()

	return user, nil
}

func (ur *UserRepository) GetAll() ([]models.User, error) {
	rows, err := DB.Query("SELECT * from users")

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

	stmt, err := DB.Prepare("SELECT id, first_name, last_name, organization, position, email, password from users WHERE id = ?")

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	user := models.User{}

	sqlErr := stmt.QueryRow(id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Organization, &user.Position, &user.Email, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, sqlErr
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (models.User, error) {

	stmt, err := DB.Prepare("SELECT id, first_name, last_name, organization, position, email, password from users WHERE email = ?")

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	user := models.User{}

	sqlErr := stmt.QueryRow(email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Organization, &user.Position, &user.Email, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, sqlErr
	}
	return user, nil
}

func (ur *UserRepository) Update(user models.User) (models.User, error) {

	tx, err := DB.Begin()

	if err != nil {
		return models.User{}, err
	}

	stmt, err := tx.Prepare("UPDATE users SET first_name = ?, last_name = ?, organization = ?, position = ?, email = ?, password = ? WHERE Id = ?")

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Organization, user.Position, user.Email, user.Password, user.Id)

	if err != nil {
		return models.User{}, err
	}

	tx.Commit()

	return user, nil
}

func (ur *UserRepository) Delete(id int) (models.User, error) {

	tx, err := DB.Begin()

	if err != nil {
		return models.User{}, err
	}

	stmt, err := DB.Prepare("DELETE from users where Id = ?")

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	user, _ := ur.GetByID(id)

	_, err = stmt.Exec(id)

	if err != nil {
		return models.User{}, err
	}

	tx.Commit()

	return user, nil
}
