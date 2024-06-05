package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserHandler struct {
	getUserUseCase use_cases.GetUserUseCase
}

func NewGetUserHandler(guuc use_cases.GetUserUseCase) GetUserHandler {
	return GetUserHandler{getUserUseCase: guuc}
}

func (guh *GetUserHandler) GetUserGin(c *gin.Context) {
	id, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return
	}

	filter := dtos.UsersFilterDTO{
		Ids: []int{int(id.(float64))},
	}
	user, err := guh.getUserUseCase.Execute(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
