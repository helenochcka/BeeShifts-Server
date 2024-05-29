package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthzHandler struct {
	loginUseCase use_cases.LoginUseCase
	authUseCase  use_cases.AuthzUseCase
}

func NewAuthzHandler(luc use_cases.LoginUseCase, auc use_cases.AuthzUseCase) AuthzHandler {
	return AuthzHandler{loginUseCase: luc, authUseCase: auc}
}

func (ah *AuthzHandler) Login(c *gin.Context) {
	regReq := dtos.LoginSchemaDTO{}
	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ah.loginUseCase.Execute(regReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (ah *AuthzHandler) AuthzUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		payload, err := ah.authUseCase.Execute(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", payload["id"])

		c.Next()
	}
}
