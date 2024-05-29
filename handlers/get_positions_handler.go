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

func (gph *GetPositionsHandler) GetPositionsGin(c *gin.Context) {
	var dto dtos.PositionsFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	positions, err := gph.getPositionsUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": positions})
}
