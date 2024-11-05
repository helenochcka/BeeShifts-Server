package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUsersHandler struct {
	getUsersUseCase use_cases.GetUsersUseCase
}

func NewGetUsersHandler(guuc use_cases.GetUsersUseCase) GetUsersHandler {
	return GetUsersHandler{getUsersUseCase: guuc}
}

// GetUsersGin	godoc
// @Summary		Get users
// @Tags		users
// @Produce		json
// @Param		id	query []int false	"User id" collectionFormat(multi)
// @Param		organization_id	query []int false	"Organization id" collectionFormat(multi)
// @Param		position_id	query []int false	"Position id" collectionFormat(multi)
// @Param		first_name	query []string false	"User's first name" collectionFormat(multi)
// @Param		last_name	query []string false	"User's last name" collectionFormat(multi)
// @Param		email	query []string false	"User's email" collectionFormat(multi)
// @Success		200  {array}  dtos.UserDTO
// @Router		/users [get]
// @Security	ApiKeyAuth
func (guh *GetUsersHandler) GetUsersGin(c *gin.Context) {
	var dto dtos.UsersFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	users, err := guh.getUsersUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
