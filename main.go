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
	organizationRepository := repositories.NewOrganizationRepo()
	positionRepository := repositories.NewPositionRepo()

	userService := services.NewUserService(userRepository, organizationRepository, positionRepository)
	getUsersUseCase := use_cases.NewGetUserUseCase(userService)
	createUserUseCase := use_cases.NewCreateUserUseCase(userService)
	userHandler := handlers.NewUserHandler(getUsersUseCase, createUserUseCase)

	//authService := services.AuthService{SecretKey: cfg.Server.SecretKey}
	//authHandler := handlers.NewAuthHandler(userService, authService)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)

	//r.POST("/login", authHandler.Login)
	//r.GET("/user", authHandler.AuthUser, userHandler.GetAllUsers)
	//r.GET("/user/:id", userHandler.GetUserById)
	//r.PUT("/user", userHandler.UpdateUser)
	//r.DELETE("/user/:id", userHandler.DeleteUser)
	//
	_ = r.Run(cfg.Server.Address + ": " + strconv.Itoa(cfg.Server.Port))
}
