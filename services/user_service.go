package services

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/entities"
	"BeeShifts-Server/repositories"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(ur repositories.UserRepo) UserService {
	return UserService{userRepo: ur}
}

func (us *UserService) GetUsers(dto dtos.UsersFilterDTO) ([]entities.UserEntity, error) {
	userEntities, err := us.userRepo.GetAll(dto)

	return userEntities, err
}

func (us *UserService) GetUser(dto dtos.UsersFilterDTO) (*entities.UserEntity, error) {
	userEntity, err := us.userRepo.GetOne(dto)
	if err != nil {
		return nil, err
	} //TODO cast repo err to domain err

	return userEntity, nil
}

func (us *UserService) CreateUser(userToCreate entities.UserEntity) (*entities.UserEntity, error) {
	createdUser, err := us.userRepo.Insert(userToCreate)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (us *UserService) UpdateUser(userToUpdate entities.UserEntity) (*entities.UserEntity, error) {
	updatedUser, err := us.userRepo.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
