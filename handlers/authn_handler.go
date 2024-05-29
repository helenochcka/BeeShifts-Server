package handlers

import (
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthnHandler struct {
	authnUseCase use_cases.AuthnUseCase
}

func NewAuthnHandler(auc use_cases.AuthnUseCase) AuthnHandler {
	return AuthnHandler{authnUseCase: auc}
}

func (ah AuthnHandler) AuthnGin(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("id")
		if exists != true {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
			return
		}

		err := ah.authnUseCase.Execute(role, int(id.(float64)))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to authenticate"})
			c.Abort()
			return
		}

		c.Next()
	}
}
