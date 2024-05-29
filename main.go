package main

import (
	"BeeShifts-Server/config"
	"BeeShifts-Server/handlers"
	"BeeShifts-Server/repositories"
	"BeeShifts-Server/services"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	cfg := config.LoadYamlConfig("config.yaml")
	_ = repositories.ConnectDatabase(cfg.DB.Host, cfg.DB.Port, cfg.DB.UserName, cfg.DB.Password, cfg.DB.DBName)
	r := gin.Default()

	userRepository := repositories.NewUserRepoPgSQL()
	organizationRepository := repositories.NewOrgRepoPgSQL()
	positionRepository := repositories.NewPositionRepoPgSQL()

	userService := services.NewUserService(userRepository)
	organizationService := services.NewOrgService(organizationRepository)
	positionService := services.NewPositionService(positionRepository)

	getUserUseCase := use_cases.NewGetUserUseCase(userService, organizationService, positionService)
	getUsersUseCase := use_cases.NewGetUsersUseCase(userService, organizationService, positionService)
	createUserUseCase := use_cases.NewCreateUserUseCase(userService)
	attachUserUseCase := use_cases.NewAttachUserUseCase(userService)
	updateUserUseCase := use_cases.NewUpdateUserUseCase(userService)
	getOrgsUseCase := use_cases.NewGetOrgsUseCase(organizationService)
	getPositionsUseCase := use_cases.NewGetPositionsUseCase(positionService)
	createPositionUseCase := use_cases.NewCreatePositionUseCase(positionService)
	updatePositionUseCase := use_cases.NewUpdatePositionUseCase(positionService)

	getUserHandler := handlers.NewGetUserHandler(getUserUseCase)
	getUsersHandler := handlers.NewGetUsersHandler(getUsersUseCase)
	createUserHandler := handlers.NewCreateUserHandler(createUserUseCase)
	attachUserHandler := handlers.NewAttachUserHandler(attachUserUseCase)
	updateUserHandler := handlers.NewUpdateUserHandler(updateUserUseCase)
	getOrgsHandler := handlers.NewGetOrgsHandler(getOrgsUseCase)
	getPositionsHandler := handlers.NewGetPositionsHandler(getPositionsUseCase)
	createPositionHandler := handlers.NewCreatePositionHandler(createPositionUseCase)
	updatePositionHandler := handlers.NewUpdatePositionHandler(updatePositionUseCase)

	authzService := services.AuthzService{SecretKey: cfg.Server.SecretKey}
	loginUseCase := use_cases.NewLoginUseCase(userService, authzService)
	authzUseCase := use_cases.NewAuthzUseCase(authzService)
	authnUseCase := use_cases.NewAuthnUseCase(userService)
	authzHandler := handlers.NewAuthzHandler(loginUseCase, authzUseCase)
	authnHandler := handlers.NewAuthnHandler(authnUseCase)

	r.POST("/sign_up", createUserHandler.CreateUserGin)
	r.POST("/login", authzHandler.Login)
	r.POST("/positions", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), createPositionHandler.CreatePositionGin)

	r.GET("/users/me", authzHandler.AuthzUser(), getUserHandler.GetUserGin)
	r.GET("/users", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), getUsersHandler.GetUsersGin)
	r.GET("/organizations", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), getOrgsHandler.GetOrgsGin)
	r.GET("/positions", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), getPositionsHandler.GetPositionsGin)

	r.PUT("/users/me", authzHandler.AuthzUser(), updateUserHandler.UpdateUserGin)
	r.PUT("/users", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), attachUserHandler.AttachUserGin)
	r.PUT("/positions", authzHandler.AuthzUser(), authnHandler.AuthnGin("Manager"), updatePositionHandler.UpdatePositionGin)

	_ = r.Run(cfg.Server.Address + ": " + strconv.Itoa(cfg.Server.Port))
}
