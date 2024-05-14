package handlers

import (
	"BeeShifts-Server/dtos"
	"BeeShifts-Server/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserSchema struct {
	Id           int
	Organization *string
	Position     *string
	FirstName    string
	LastName     string
	Email        string
}

type UserHandler struct {
	getUsersUseCase   use_cases.GetUserUseCase
	createUserUseCase use_cases.CreateUserUseCase
}

func NewUserHandler(guuc use_cases.GetUserUseCase, cuuc use_cases.CreateUserUseCase) UserHandler {
	return UserHandler{getUsersUseCase: guuc, createUserUseCase: cuuc}
}

func (uh *UserHandler) GetUsers(c *gin.Context) {

	var dto dtos.UsersFilterDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	users, err := uh.getUsersUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
	}

	var userSchemas []UserSchema
	for _, user := range users {
		userSchema := uh.dtoToSchema(user)
		userSchemas = append(userSchemas, userSchema)
	}

	c.JSON(http.StatusOK, gin.H{"data": userSchemas})
}

func (uh *UserHandler) CreateUser(c *gin.Context) {

	var dto dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := uh.createUserUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": uh.dtoToSchema(*user)})
}

func (uh *UserHandler) dtoToSchema(dto dtos.UserDTO) UserSchema {
	schema := UserSchema{
		Id:           dto.Id,
		Organization: dto.Organization,
		Position:     dto.Position,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Email:        dto.Email,
	}
	return schema
}

//func (uh *UserHandler) UpdateUser(c *gin.Context) {
//
//	var json models.User
//
//	if err := c.ShouldBindJSON(&json); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	user, err := uh.userService.UpdateUser(json)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	c.JSON(http.StatusOK, gin.H{"data": user})
//}
//
//func (uh *UserHandler) DeleteUser(c *gin.Context) {
//
//	id, err := strconv.Atoi(c.Param("id"))
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//	}
//
//	err = uh.userService.DeleteUser(id)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	c.JSON(http.StatusOK, gin.H{})
//}
