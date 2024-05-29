package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AttachUserHandler struct {
	attachUserUseCase use_cases.AttachUserUseCase
}

func NewAttachUserHandler(auuc use_cases.AttachUserUseCase) AttachUserHandler {
	return AttachUserHandler{attachUserUseCase: auuc}
}

func (auh *AttachUserHandler) AttachUserGin(c *gin.Context) {
	var dto dtos.AttachUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := auh.attachUserUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
