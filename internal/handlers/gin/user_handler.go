package gin

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlerGin struct {
	getUserUseCase    usecases.GetUserUseCase
	getUsersUseCase   usecases.GetUsersUseCase
	createUserUseCase usecases.CreateUserUseCase
	attachUserUseCase usecases.AttachUserUseCase
	updateUserUseCase usecases.UpdateUserUseCase
}

func NewUserHandlerGin(
	guuc usecases.GetUserUseCase,
	gusuc usecases.GetUsersUseCase,
	cuuc usecases.CreateUserUseCase,
	auuc usecases.AttachUserUseCase,
	uuuc usecases.UpdateUserUseCase) UserHandlerGin {

	return UserHandlerGin{
		getUserUseCase:    guuc,
		getUsersUseCase:   gusuc,
		createUserUseCase: cuuc,
		attachUserUseCase: auuc,
		updateUserUseCase: uuuc}
}

// GetOne	godoc
// @Summary		Get current user
// @Tags		users
// @Produce		json
// @Success		200	{object}	users.ViewDTO
// @Router		/users/me [get]
// @Security	ApiKeyAuth
func (uhg *UserHandlerGin) GetOne(c *gin.Context) {
	id, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
		return
	}

	filter := users.FilterDTO{
		Ids: []int{id.(int)},
	}
	userDTO, err := uhg.getUserUseCase.Execute(filter)

	if err != nil {
		uhg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userDTO})
}

// GetMany	godoc
// @Summary		Get users
// @Tags		users
// @Produce		json
// @Param		id	query []int false	"User id" collectionFormat(multi)
// @Param		organization_id	query []int false	"Organization id" collectionFormat(multi)
// @Param		position_id	query []int false	"Position id" collectionFormat(multi)
// @Param		first_name	query []string false	"User's first name" collectionFormat(multi)
// @Param		last_name	query []string false	"User's last name" collectionFormat(multi)
// @Param		email	query []string false	"User's email" collectionFormat(multi)
// @Success		200  {array}	users.ViewDTO
// @Router		/users [get]
// @Security	ApiKeyAuth
func (uhg *UserHandlerGin) GetMany(c *gin.Context) {
	var filter users.FilterDTO

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query arguments, " + err.Error()})
		return
	}

	userDTOS, err := uhg.getUsersUseCase.Execute(filter)

	if err != nil {
		uhg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userDTOS})
}

// Create	godoc
// @Summary			Create new user
// @Tags			users
// @Produce			json
// @Param			User	body		users.CreateDTO	true	"User JSON"
// @Success			201		{object}	users.ViewDTO
// @Router			/sign_up [post]
func (uhg *UserHandlerGin) Create(c *gin.Context) {
	var dto users.CreateDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	userEntity, err := uhg.createUserUseCase.Execute(dto)

	if err != nil {
		uhg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": userEntity})
}

// Attach	godoc
// @Summary			Attach user to organization and set position
// @Tags			users
// @Produce			json
// @Param			AttachUserInfo	body		users.AttachDTO	true	"Data for users attachment JSON"
// @Success			200				{object}	users.Entity
// @Router			/users [put]
// @Security		ApiKeyAuth
func (uhg *UserHandlerGin) Attach(c *gin.Context) {
	var dto users.AttachDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	userEntity, err := uhg.attachUserUseCase.Execute(dto)

	if err != nil {
		uhg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userEntity})
}

// Update	godoc
// @Summary			Update user
// @Tags			users
// @Produce			json
// @Param			User	body		users.UpdateSelfDTO	true	"User to update JSON"
// @Success			200		{object}	users.Entity
// @Router			/users/me [put]
// @Security		ApiKeyAuth
func (uhg *UserHandlerGin) Update(c *gin.Context) {
	id, exists := c.Get("id")
	if exists != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
		return
	}

	var dto users.UpdateSelfDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	userEntity, err := uhg.updateUserUseCase.Execute(id.(int), dto)

	if err != nil {
		uhg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userEntity})
}

func (uhg *UserHandlerGin) mapUsersErrToHTTPErr(err error, c *gin.Context) {
	switch {
	case errors.Is(err, users.EmailAlreadyUsed):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, users.UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, users.RoleDoesNotExist):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
