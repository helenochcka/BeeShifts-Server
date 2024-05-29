package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateUserHandler struct {
	updateUserUseCase use_cases.UpdateUserUseCase
}

func NewUpdateUserHandler(uuuc use_cases.UpdateUserUseCase) UpdateUserHandler {
	return UpdateUserHandler{updateUserUseCase: uuuc}
}

func (uuh *UpdateUserHandler) UpdateUserGin(c *gin.Context) {
	id, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return
	}

	var dto dtos.UpdateSelfUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := uuh.updateUserUseCase.Execute(int(id.(float64)), dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
