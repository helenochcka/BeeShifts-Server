package gin

import (
	"BeeShifts-Server/internal/core/users"
	"BeeShifts-Server/internal/core/users/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlerGin struct {
	loginUseCase        usecases.LoginUseCase
	authenticateUseCase usecases.AuthenticateUseCase
	authorizeUseCase    usecases.AuthorizeUseCase
}

func NewAuthHandlerGin(
	luc usecases.LoginUseCase,
	anuc usecases.AuthenticateUseCase,
	azuc usecases.AuthorizeUseCase,
) AuthHandlerGin {

	return AuthHandlerGin{
		loginUseCase:        luc,
		authenticateUseCase: anuc,
		authorizeUseCase:    azuc,
	}
}

// Login			godoc
// @Summary			Get API token
// @Description		Returns API token by credentials (password should be hashed).
// @Tags			users
// @Produce			json
// @Param			CredsDTO	body		users.CredsDTO	true	"CredsDTO JSON"
// @Success			200			{object}	users.LoginResponseDTO
// @Router			/login [post]
func (ahg *AuthHandlerGin) Login(c *gin.Context) {
	creds := users.CredsDTO{}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body, " + err.Error()})
		return
	}

	token, err := ahg.loginUseCase.Execute(creds)
	if err != nil {
		ahg.mapUsersErrToHTTPErr(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (ahg *AuthHandlerGin) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		id, err := ahg.authenticateUseCase.Execute(token)

		if err != nil {
			ahg.mapUsersErrToHTTPErr(err, c)
			c.Abort()
			return
		}

		c.Set("id", *id)

		c.Next()
	}
}

func (ahg *AuthHandlerGin) AuthorizeGin(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("id")
		if exists != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is missing in request context"})
			return
		}

		err := ahg.authorizeUseCase.Execute(role, id.(int))

		if err != nil {
			ahg.mapUsersErrToHTTPErr(err, c)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (ahg *AuthHandlerGin) mapUsersErrToHTTPErr(err error, c *gin.Context) {
	switch {
	case errors.Is(err, users.MultipleUsersFound):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, users.UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, users.IncorrectCredentials) ||
		errors.Is(err, users.TokenExpired) ||
		errors.Is(err, users.TokenSignatureInvalid):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case errors.Is(err, users.InsufficientRights):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
