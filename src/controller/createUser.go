package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/controller/dto"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	//passando interfaces ao inves de objetos
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")

}
