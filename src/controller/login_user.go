package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	dto_request "github.com/vinialeixo/crud-golang/src/controller/dto/request"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {

	logger.Info("Init loginUser controller", zap.String("journey", "loginUser"))

	var userRequest dto_request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info ", err, zap.String("journey", "loginUser")) //string message
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)
	//passando interfaces ao inves de objetos

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "loginUser"),
		zap.String("userId", domain.GetID()),
	)
	fmt.Println(token)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))

}
