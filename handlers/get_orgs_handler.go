package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOrgsHandler struct {
	getOrgsUseCase use_cases.GetOrgsUseCase
}

func NewGetOrgsHandler(gouc use_cases.GetOrgsUseCase) GetOrgsHandler {
	return GetOrgsHandler{getOrgsUseCase: gouc}
}

// GetOrgsGin	godoc
// @Summary		Get organizations
// @Tags		organizations
// @Produce		json
// @Param		id		query	[]int		false	"Organization id"	collectionFormat(multi)
// @Param		name	query	[]string	false	"Organization name"	collectionFormat(multi)
// @Success		200		{array}	entities.OrganizationEntity
// @Router		/organizations [get]
// @Security	ApiKeyAuth
func (goh *GetOrgsHandler) GetOrgsGin(c *gin.Context) {
	var dto dtos.OrgsFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	organizations, err := goh.getOrgsUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": organizations})
}
