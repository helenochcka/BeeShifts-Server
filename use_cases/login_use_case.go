package use_cases

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/services"
	"errors"
)

type LoginUseCase struct {
	userService  services.UserService
	authzService services.AuthzService
}

func NewLoginUseCase(us services.UserService, as services.AuthzService) LoginUseCase {
	return LoginUseCase{userService: us, authzService: as}
}

func (luc *LoginUseCase) Execute(regReq dtos.LoginSchemaDTO) (*dtos.LoginResponse, error) {
	userFilter := dtos.UsersFilterDTO{
		Emails: []string{regReq.Email},
	}

	user, err := luc.userService.GetUser(userFilter)

	if err != nil {
		return nil, err
	}

	if user.Password != regReq.Password {
		return nil, errors.New("incorrect password")
	}

	token, err := luc.authzService.GenerateToken(user.Id)

	if err != nil {
		return nil, err
	}
	return &dtos.LoginResponse{AccessToken: token}, nil
}
