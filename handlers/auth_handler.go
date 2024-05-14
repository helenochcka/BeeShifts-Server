package handlers

import (
	"BeeShifts-Server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthHandler struct {
	userService services.UserService
	authService services.AuthService
}

func NewAuthHandler(us services.UserService, as services.AuthService) AuthHandler {
	return AuthHandler{userService: us, authService: as}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	regReq := LoginSchema{}
	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//user, err := ah.userService.GetUserByEmail(regReq.Email)

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//if user.Password != regReq.Password {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//token, err := ah.authService.GenerateToken(user.Id)
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"data": LoginResponse{AccessToken: "blabla"}})
}

func (ah *AuthHandler) AuthUser(c *gin.Context) {
	payload, err := ah.authService.PayloadFromToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Set("user_id", payload["id"])

	c.Next()
}
