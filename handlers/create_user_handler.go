package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserHandler struct {
	createUserUseCase use_cases.CreateUserUseCase
}

func NewCreateUserHandler(cuuc use_cases.CreateUserUseCase) CreateUserHandler {
	return CreateUserHandler{createUserUseCase: cuuc}
}

func (cuh *CreateUserHandler) CreateUserGin(c *gin.Context) {
	var dto dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := cuh.createUserUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
