package main

import (
	"BeeShifts-Server/config"
	"BeeShifts-Server/handlers"
	"BeeShifts-Server/repositories"
	"BeeShifts-Server/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	cfg := config.LoadYamlConfig("config.yaml")
	_ = repositories.ConnectDatabase(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)
	r := gin.Default()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	authService := services.AuthService{SecretKey: cfg.Server.SecretKey}
	authHandler := handlers.NewAuthHandler(userService, authService)

	r.POST("/login", authHandler.Login)
	r.POST("/user", userHandler.CreateUser)
	r.GET("/user", authHandler.AuthUser, userHandler.GetAllUsers)
	r.GET("/user/:id", userHandler.GetUserById)
	r.PUT("/user", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)

	_ = r.Run(cfg.Server.Address + ": " + strconv.Itoa(cfg.Server.Port))
}
