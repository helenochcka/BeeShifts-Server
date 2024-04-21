package handlers

import (
	"BeeShifts-Server/models"
	"BeeShifts-Server/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return UserHandler{userService: us}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.CreateUser(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, _ := uh.userService.GetUsers()

	if users == nil {
		c.JSON(http.StatusOK, gin.H{"error": "No Records Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uh *UserHandler) GetUserById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, _ := uh.userService.GetUserByID(id)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {

	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.UpdateUser(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	user, err := uh.userService.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
