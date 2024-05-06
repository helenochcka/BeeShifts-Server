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
	users, err := us.userRepository.GetAll()

	return users, err
}

func (us *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := us.userRepository.GetByID(id)

	if errors.Is(err, repositories.RecNotFound) {
		return nil, UserNotFound
	}

	return user, err
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.userRepository.GetByEmail(email)

	if errors.Is(err, repositories.RecNotFound) {
		return nil, UserNotFound
	}

	return user, err
}

func (us *UserService) UpdateUser(user models.User) (*models.User, error) {
	updUser, err := us.userRepository.Update(user)

	return updUser, err
}

func (us *UserService) DeleteUser(id int) (*models.User, error) {
	delUser, err := us.userRepository.Delete(id)

	if errors.Is(err, repositories.RecNotFound) {
		return nil, UserNotFound
	}

	return delUser, err
}
