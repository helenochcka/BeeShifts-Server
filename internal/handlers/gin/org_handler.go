package gin

import (
	"BeeShifts-Server/internal/core/organizations"
	"BeeShifts-Server/internal/core/organizations/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrgHandlerGin struct {
	getOrgsUseCase usecases.GetOrgsUseCase
}

func NewOrgHandlerGin(gouc usecases.GetOrgsUseCase) OrgHandlerGin {
	return OrgHandlerGin{getOrgsUseCase: gouc}
}

// GetMany	godoc
// @Summary		Get organizations
// @Tags		organizations
// @Produce		json
// @Param		id		query	[]int		false	"Organization id"	collectionFormat(multi)
// @Param		name	query	[]string	false	"Organization name"	collectionFormat(multi)
// @Success		200		{array}	organizations.Entity
// @Router		/organizations [get]
// @Security	ApiKeyAuth
func (ohg *OrgHandlerGin) GetMany(c *gin.Context) {
	var filter organizations.FilterDTO

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query arguments, " + err.Error()})
		return
	}

	organizationEntities, err := ohg.getOrgsUseCase.Execute(filter)

	if err != nil {
		ohg.mapOrgsErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": organizationEntities})
}

func (ohg *OrgHandlerGin) mapOrgsErrToHTTPErr(err error, c *gin.Context) {
	switch {
	case errors.Is(err, organizations.MultipleOrgsFound):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, organizations.OrgNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
