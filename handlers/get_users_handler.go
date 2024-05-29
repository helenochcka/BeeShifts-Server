package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUsersHandler struct {
	getUsersUseCase use_cases.GetUsersUseCase
}

func NewGetUsersHandler(guuc use_cases.GetUsersUseCase) GetUsersHandler {
	return GetUsersHandler{getUsersUseCase: guuc}
}

func (guh *GetUsersHandler) GetUsersGin(c *gin.Context) {
	var dto dtos.UsersFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	users, err := guh.getUsersUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
