package gin

import (
	"BeeShifts-Server/internal/core/positions"
	"BeeShifts-Server/internal/core/positions/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PositionHandlerGin struct {
	getPositionsUseCase   usecases.GetPositionsUseCase
	updatePositionUseCase usecases.UpdatePositionUseCase
	createPositionUseCase usecases.CreatePositionUseCase
}

func NewPositionHandlerGin(
	gpuc usecases.GetPositionsUseCase,
	upuc usecases.UpdatePositionUseCase,
	cpuc usecases.CreatePositionUseCase) PositionHandlerGin {

	return PositionHandlerGin{
		getPositionsUseCase:   gpuc,
		updatePositionUseCase: upuc,
		createPositionUseCase: cpuc}
}

// GetMany	godoc
// @Summary			Get positions
// @Tags			positions
// @Produce			json
// @Param			id			query	[]int		false	"Position id"	collectionFormat(multi)
// @Param			manager_id	query	[]int		false	"Manager id"	collectionFormat(multi)
// @Param			name		query	[]string	false	"Position name"	collectionFormat(multi)
// @Success			200			{array}	positions.Entity
// @Router			/positions [get]
// @Security		ApiKeyAuth
func (phg *PositionHandlerGin) GetMany(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
		return
	}
	var filter positions.FilterDTO

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query arguments, " + err.Error()})
		return
	}

	positionEntities, err := phg.getPositionsUseCase.Execute(managerId.(int), filter)

	if err != nil {
		phg.mapPositionsErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": positionEntities})
}

// Update	godoc
// @Summary				Update position
// @Tags				positions
// @Produce				json
// @Param				Position	body		positions.UpdateDTO	true	"Position to update JSON"
// @Success				200			{object}	positions.Entity
// @Router 				/positions [put]
// @Security			ApiKeyAuth
func (phg *PositionHandlerGin) Update(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
		return
	}

	var dto positions.UpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	positionEntity, err := phg.updatePositionUseCase.Execute(managerId.(int), dto)

	if err != nil {
		phg.mapPositionsErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": positionEntity})
}

// Create	godoc
// @Summary				Create new position
// @Tags				positions
// @Produce				json
// @Param				Position	body		positions.CreateDTO	true	"Position JSON"
// @Success				201			{object}	positions.Entity
// @Router				/positions [post]
// @Security			ApiKeyAuth
func (phg *PositionHandlerGin) Create(c *gin.Context) {
	managerId, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
		return
	}

	var dto positions.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	positionEntity, err := phg.createPositionUseCase.Execute(managerId.(int), dto)

	if err != nil {
		phg.mapPositionsErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": positionEntity})
}

func (phg *PositionHandlerGin) mapPositionsErrToHTTPErr(err error, c *gin.Context) {
	switch {
	case errors.Is(err, positions.PositionNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
