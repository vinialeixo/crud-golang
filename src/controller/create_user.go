package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	dto_request "github.com/vinialeixo/crud-golang/src/controller/dto/request"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest dto_request.UserRequest

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

	domainResult, errs := uc.service.CreateUser(domain)
	if errs != nil {
		c.JSON(errs.Code, errs)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
		zap.String("username", domain.GetName()),
		zap.String("email", domain.GetEmail()),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))

}
