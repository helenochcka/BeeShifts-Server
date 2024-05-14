package services

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/repositories"
	"BeeShifts-Server/repositories/models"
)

type UserService struct {
	userRepo         repositories.UserRepo
	organizationRepo repositories.OrganizationRepo
	positionRepo     repositories.PositionRepo
}

func NewUserService(ur repositories.UserRepo, or repositories.OrganizationRepo, pr repositories.PositionRepo) UserService {
	return UserService{userRepo: ur, organizationRepo: or, positionRepo: pr}
}

func (us *UserService) GetUsers(dto dtos.UsersFilterDTO) ([]dtos.UserDTO, error) {

	userModels, err := us.userRepo.GetAll(dto)
	if err != nil {
		return nil, err
	}

	var userDTOS []dtos.UserDTO
	for _, userModel := range userModels {
		userDTO, err := us.generateUserDTO(userModel)
		if err != nil {
			return nil, err
		}
		userDTOS = append(userDTOS, *userDTO)
	}

	return userDTOS, nil
}

func (us *UserService) GetUser(dto dtos.UsersFilterDTO) (*dtos.UserDTO, error) {
	userModel, err := us.userRepo.GetOne(dto)
	if err != nil {
		return nil, err
	}

	userDTO, err := us.generateUserDTO(*userModel)
	if err != nil {
		return nil, err
	}

	return userDTO, nil
}

func (us *UserService) CreateUser(dto dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	userToCreate := models.User{
		OrganizationId: nil,
		PositionId:     nil,
		Role:           dto.Role,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
	}

	createdUser, err := us.userRepo.Insert(userToCreate)
	if err != nil {
		return nil, err
	}

	userDTO, err := us.generateUserDTO(*createdUser)
	if err != nil {
		return nil, err
	}

	return userDTO, err
}

func (us *UserService) UpdateUser(dto dtos.UpdateUserDTO) (*dtos.UserDTO, error) {
	userToUpdate := models.User{
		Id:             dto.Id,
		OrganizationId: dto.Organization,
		PositionId:     dto.Position,
		Role:           dto.Role,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
	}
	updatedUser, err := us.userRepo.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	userDTO, err := us.generateUserDTO(*updatedUser)
	if err != nil {
		return nil, err
	}

	return userDTO, err
}

func (us *UserService) generateUserDTO(userModel models.User) (*dtos.UserDTO, error) {
	organizationName, err := us.getOrganizationNameById(userModel.OrganizationId)
	if err != nil {
		return nil, err
	}
	positionName, err := us.getPositionNameById(userModel.PositionId)
	if err != nil {
		return nil, err
	}
	userDTO := dtos.UserDTO{
		Id:           userModel.Id,
		Organization: organizationName,
		Position:     positionName,
		Role:         userModel.Role,
		FirstName:    userModel.FirstName,
		LastName:     userModel.LastName,
		Email:        userModel.Email,
		Password:     userModel.Password,
	}

	return &userDTO, nil
}

func (us *UserService) getOrganizationNameById(organizationId *int) (*string, error) {
	if organizationId == nil {
		return nil, nil
	}
	organizationDTO := dtos.GetOrganizationsDTO{Ids: []int{*organizationId}}
	organization, err := us.organizationRepo.GetOne(organizationDTO)
	if err != nil {
		return nil, err
	}

	return &organization.Name, nil
}

func (us *UserService) getPositionNameById(positionId *int) (*string, error) {
	if positionId == nil {
		return nil, nil
	}
	positionDTO := dtos.GetPositionsDTO{Ids: []int{*positionId}}
	position, err := us.positionRepo.GetOne(positionDTO)
	if err != nil {
		return nil, err
	}

	return &position.Name, nil
}
