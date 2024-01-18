package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	dto_request "github.com/vinialeixo/crud-golang/src/controller/dto/request"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))
	var userRequest dto_request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info ", err, zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		rest_err.NewBadRequestError("Invalid userId,must be a hex value")
	}

	domain := model.NewUpdateUserDomain(userRequest.Name, userRequest.Age)
	//passando interfaces ao inves de objetos

	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("journey", "updateUser"),
		zap.String("username", domain.GetName()),
		zap.String("userId", domain.GetID()),
	)

	c.Status(http.StatusOK)
}
