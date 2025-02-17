package main

import (
	"BeeShifts-Server/config"
	_ "BeeShifts-Server/docs"
	orgsServices "BeeShifts-Server/internal/core/organizations/services"
	orgsUsecases "BeeShifts-Server/internal/core/organizations/usecases"
	positionsServices "BeeShifts-Server/internal/core/positions/services"
	positionsUsecases "BeeShifts-Server/internal/core/positions/usecases"
	userRoles "BeeShifts-Server/internal/core/users"
	usersServices "BeeShifts-Server/internal/core/users/services"
	usersUsecases "BeeShifts-Server/internal/core/users/usecases"
	handlersGin "BeeShifts-Server/internal/handlers/gin"
	"BeeShifts-Server/internal/middlewares"
	"BeeShifts-Server/internal/repositories/postgres"
	"BeeShifts-Server/pkg/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"os"
	"strconv"
)

// @title BeeShifts-Server API
// @version 1.0
// @securitydefinitions.apikey ApiKeyAuth
// @in	header
// @name	Authorization
// @query.collection.format multi
// @host      localhost:8000
func main() {

	cfg := config.LoadYamlConfig("config/config.yaml")
	_ = db.ConnectDatabase(cfg.DB.Host, cfg.DB.Port, cfg.DB.UserName, cfg.DB.Password, cfg.DB.DBName)
	r := gin.Default()

	var slogHandler slog.Handler = slog.NewJSONHandler(os.Stdout, nil)

	if cfg.Environment == "dev" {
		slogHandler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(slogHandler)

	userRepository := postgres.NewUserRepoPgSQL()
	organizationRepository := postgres.NewOrgRepoPgSQL()
	positionRepository := postgres.NewPositionRepoPgSQL()

	userService := usersServices.NewUserService(userRepository)
	organizationService := orgsServices.NewOrgService(organizationRepository)
	positionService := positionsServices.NewPositionService(positionRepository)

	getUserUseCase := usersUsecases.NewGetUserUseCase(userService, organizationService, positionService)
	getUsersUseCase := usersUsecases.NewGetUsersUseCase(userService, organizationService, positionService)
	createUserUseCase := usersUsecases.NewCreateUserUseCase(userService)
	attachUserUseCase := usersUsecases.NewAttachUserUseCase(userService, positionService)
	detachUserUseCase := usersUsecases.NewDetachUserUseCase(userService)
	updateUserUseCase := usersUsecases.NewUpdateUserUseCase(userService)
	getOrgsUseCase := orgsUsecases.NewGetOrgsUseCase(organizationService)
	getPositionsUseCase := positionsUsecases.NewGetPositionsUseCase(positionService)
	createPositionUseCase := positionsUsecases.NewCreatePositionUseCase(positionService)
	updatePositionUseCase := positionsUsecases.NewUpdatePositionUseCase(positionService)

	userHandler := handlersGin.NewUserHandlerGin(getUserUseCase, getUsersUseCase, createUserUseCase, attachUserUseCase, detachUserUseCase, updateUserUseCase)
	orgHandler := handlersGin.NewOrgHandlerGin(getOrgsUseCase)
	positionHandler := handlersGin.NewPositionHandlerGin(getPositionsUseCase, updatePositionUseCase, createPositionUseCase)

	authService := usersServices.AuthService{SecretKey: cfg.Server.SecretKey, TokenExpTime: cfg.Server.TokenExpTime}
	loginUseCase := usersUsecases.NewLoginUseCase(userService, authService)
	authenticateUseCase := usersUsecases.NewAuthenticateUseCase(authService)
	authorizeUseCase := usersUsecases.NewAuthorizeUseCase(userService)
	authHandler := handlersGin.NewAuthHandlerGin(loginUseCase, authenticateUseCase, authorizeUseCase)

	r.Use(middlewares.RequestId(), middlewares.Logging(logger))

	r.POST("/sign_up", userHandler.Create)
	r.POST("/login", authHandler.Login)
	r.POST("/positions", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), positionHandler.Create)

	r.GET("/users/me", authHandler.AuthenticateUser(), userHandler.GetOne)
	r.GET("/users", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), userHandler.GetMany)
	r.GET("/organizations", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), orgHandler.GetMany)
	r.GET("/positions", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), positionHandler.GetMany)

	r.PUT("/users/me", authHandler.AuthenticateUser(), userHandler.Update)
	r.PUT("/users/attach", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), userHandler.Attach)
	r.PUT("/users/detach", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), userHandler.Detach)
	r.PUT("/positions", authHandler.AuthenticateUser(), authHandler.AuthorizeGin(userRoles.Manager), positionHandler.Update)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Run(cfg.Server.Address + ": " + strconv.Itoa(cfg.Server.Port))
}
