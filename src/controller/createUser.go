package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/controller/dto"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest dto.UserRequest

	//&userRequest To be able to modify the original userRequest variable (defined outside the function) with the data from the request,
	//you need to pass a pointer to that variable
	//By using &userRequest, you are passing the memory address of the userRequest variable to the ShouldBind function. This allows ShouldBind to modify the actual variable's content,
	err := c.ShouldBind(&userRequest)
	if err != nil {
		logger.Error("Error trying to validate user info ", err, zap.String("journey", "createUser")) //string message
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := dto.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, response)
}
