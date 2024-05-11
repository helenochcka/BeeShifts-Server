package main

import (
	"BeeShifts-Server/config"
	"BeeShifts-Server/repositories"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadYamlConfig("config.yaml")
	_ = repositories.ConnectDatabase(cfg.DB.Host, cfg.DB.Port, cfg.DB.UserName, cfg.DB.Password, cfg.DB.DBName)
	//r := gin.Default()

	userRepository := repositories.NewUserRepository()

	filter := repositories.UserFilter{FirstNames: []interface{}{"Kirill"}, LastNames: []interface{}{"Osa"}}

	users, err := userRepository.Get(filter)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Organization: %d, Position: %d, Name: %s %s\n", user.Id, user.Organization, user.Position, user.FirstName, user.LastName)
	}

	//userService := services.NewUserService(userRepository)
	//userHandler := handlers.NewUserHandler(userService)
	//
	//authService := services.AuthService{SecretKey: cfg.Server.SecretKey}
	//authHandler := handlers.NewAuthHandler(userService, authService)
	//
	//r.POST("/login", authHandler.Login)
	//r.POST("/user", userHandler.CreateUser)
	//r.GET("/user", authHandler.AuthUser, userHandler.GetAllUsers)
	//r.GET("/user/:id", userHandler.GetUserById)
	//r.PUT("/user", userHandler.UpdateUser)
	//r.DELETE("/user/:id", userHandler.DeleteUser)
	//
	//_ = r.Run(cfg.Server.Address + ": " + strconv.Itoa(cfg.Server.Port))
}
