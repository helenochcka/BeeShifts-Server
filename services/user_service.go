package services

import (
	"BeeShifts-Server/models"
	"BeeShifts-Server/repositories"
	"errors"
	_ "strconv"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return UserService{userRepository: ur}
}

func (us *UserService) CreateUser(user models.User) (*models.User, error) {
	newUser, err := us.userRepository.Add(user)

	return newUser, err
}

func (us *UserService) GetUsers() ([]models.User, error) {
	users, err := us.userRepository.Get(repositories.UserFilter{})

	//filter := repositories.UserFilter{FirstNames: []interface{}{"Kirill"}, LastNames: []interface{}{"Ponam"}}
	//
	//users, err := userRepository.GetMany(filter)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, user := range users {
	//	fmt.Printf("ID: %d, Organization: %d, Position: %d, Name: %s %s\n", user.Id, user.Organization, user.Position, user.FirstName, user.LastName)
	//}

	return users, err
}

func (us *UserService) GetUsersByPosition(position_id int) ([]models.User, error) {
	users, err := us.userRepository.Get(repositories.UserFilter{})

	return users, err
}

func (us *UserService) GetUsersByOrganization(organization_id int) ([]models.User, error) {
	users, err := us.userRepository.Get(repositories.UserFilter{})

	return users, err
}

func (us *UserService) GetUserByID(id int) ([]models.User, error) {
	user, err := us.userRepository.Get(repositories.UserFilter{})

	if errors.Is(err, repositories.RecNotFound) {
		return nil, UserNotFound
	}

	return user, err
}

func (us *UserService) GetUserByEmail(email string) ([]models.User, error) {
	user, err := us.userRepository.Get(repositories.UserFilter{})

	if errors.Is(err, repositories.RecNotFound) {
		return nil, UserNotFound
	}

	return user, err
}

func (us *UserService) UpdateUser(user models.User) (*models.User, error) {
	updUser, err := us.userRepository.Update(user)

	return updUser, err
}

func (us *UserService) DeleteUser(id int) error {
	err := us.userRepository.Delete(id)

	if errors.Is(err, repositories.RecNotFound) {
		return UserNotFound
	}

	return err
}
