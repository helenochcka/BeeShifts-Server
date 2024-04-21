package services

import (
	"BeeShifts-Server/models"
	"BeeShifts-Server/repositories"
	_ "strconv"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return UserService{userRepository: ur}
}

func (us *UserService) CreateUser(user models.User) (models.User, error) {
	user, err := us.userRepository.Add(user)

	return user, err
}

func (us *UserService) GetUsers() ([]models.User, error) {
	users, err := us.userRepository.GetAll()

	return users, err
}

func (us *UserService) GetUserByID(id int) (models.User, error) {
	user, err := us.userRepository.GetByID(id)

	return user, err
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	user, err := us.userRepository.GetByEmail(email)

	return user, err
}

func (us *UserService) UpdateUser(user models.User) (models.User, error) {
	user, err := us.userRepository.Update(user)

	return user, err
}

func (us *UserService) DeleteUser(id int) (models.User, error) {
	user, err := us.userRepository.Delete(id)

	return user, err
}
