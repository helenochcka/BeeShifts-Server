package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetPositionsHandler struct {
	getPositionsUseCase use_cases.GetPositionsUseCase
}

func NewGetPositionsHandler(gpuc use_cases.GetPositionsUseCase) GetPositionsHandler {
	return GetPositionsHandler{getPositionsUseCase: gpuc}
}

// GetPositionsGin	godoc
// @Summary			Get positions
// @Tags			positions
// @Produce			json
// @Param			id			query	[]int		false	"Position id"	collectionFormat(multi)
// @Param			manager_id	query	[]int		false	"Manager id"	collectionFormat(multi)
// @Param			name		query	[]string	false	"Position name"	collectionFormat(multi)
// @Success			200			{array}	entities.PositionEntity
// @Router			/positions [get]
// @Security		ApiKeyAuth
func (gph *GetPositionsHandler) GetPositionsGin(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return
	}
	var dto dtos.PositionsFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	positions, err := gph.getPositionsUseCase.Execute(int(managerId.(float64)), dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": positions})
}
