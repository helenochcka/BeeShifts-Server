package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePositionHandler struct {
	updatePositionUseCase use_cases.UpdatePositionUseCase
}

func NewUpdatePositionHandler(upuc use_cases.UpdatePositionUseCase) UpdatePositionHandler {
	return UpdatePositionHandler{updatePositionUseCase: upuc}
}

// UpdatePositionGin	godoc
// @Summary				Update position
// @Tags				positions
// @Produce				json
// @Param				Position	body		dtos.UpdatePositionDTO	true	"Position to update JSON"
// @Success				200			{object}	entities.PositionEntity
// @Router 				/positions [put]
// @Security			ApiKeyAuth
func (uph *UpdatePositionHandler) UpdatePositionGin(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return
	}

	var dto dtos.UpdatePositionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	position, err := uph.updatePositionUseCase.Execute(int(managerId.(float64)), dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": position})
}
