package services

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/repositories"
	"BeeShifts-Server/internal/repositories/ifaces"
	"errors"
)

type UserService struct {
	userRepo ifaces.UserRepo
}

func NewUserService(ur ifaces.UserRepo) UserService {
	return UserService{userRepo: ur}
}

func (us *UserService) GetUsers(filter users.FilterDTO) ([]users.Entity, error) {
	userEntities, err := us.userRepo.GetAll(filter)
	if err != nil {
		return nil, us.mapRepoErrToUsersErr(err)
	}
	return userEntities, nil
}

func (us *UserService) GetUser(filter users.FilterDTO) (*users.Entity, error) {
	userEntity, err := us.userRepo.GetOne(filter)
	if err != nil {
		return nil, us.mapRepoErrToUsersErr(err)
	}

	return userEntity, nil
}

func (us *UserService) FindUser(filter users.FilterDTO) (*users.Entity, error) {
	userEntity, err := us.userRepo.GetOneOrNil(filter)
	if err != nil {
		return nil, us.mapRepoErrToUsersErr(err)
	}

	return userEntity, nil
}

func (us *UserService) CreateUser(userToCreate users.Entity) (*users.Entity, error) {
	createdUser, err := us.userRepo.Insert(userToCreate)
	if err != nil {
		return nil, us.mapRepoErrToUsersErr(err)
	}
	return createdUser, nil
}

func (us *UserService) UpdateUser(userToUpdate users.Entity) (*users.Entity, error) {
	updatedUser, err := us.userRepo.Update(userToUpdate)
	if err != nil {
		return nil, us.mapRepoErrToUsersErr(err)
	}

	return updatedUser, nil
}

func (us *UserService) IsRoleValid(role string) bool {
	if _, ok := users.Roles[role]; !ok {
		return false
	}
	return true
}

func (us *UserService) mapRepoErrToUsersErr(err error) error {
	switch {
	case errors.Is(err, repositories.MultipleRecFound):
		return users.MultipleUsersFound
	case errors.Is(err, repositories.RecNotFound):
		return users.UserNotFound
	default:
		return err
	}
}
