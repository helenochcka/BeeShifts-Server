package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePositionHandler struct {
	createPositionUseCase use_cases.CreatePositionUseCase
}

func NewCreatePositionHandler(cpuc use_cases.CreatePositionUseCase) CreatePositionHandler {
	return CreatePositionHandler{createPositionUseCase: cpuc}
}

// CreatePositionGin	godoc
// @Summary				Create new position
// @Tags				positions
// @Produce				json
// @Param				Position	body		dtos.CreatePositionDTO	true	"Position JSON"
// @Success				201			{object}	entities.PositionEntity
// @Router				/positions [post]
// @Security			ApiKeyAuth
func (cph *CreatePositionHandler) CreatePositionGin(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return
	}

	var dto dtos.CreatePositionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	position, err := cph.createPositionUseCase.Execute(int(managerId.(float64)), dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": position})
}
