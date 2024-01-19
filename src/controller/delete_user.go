package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	dto_request "github.com/vinialeixo/crud-golang/src/controller/dto/request"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))
	var userRequest dto_request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info ", err, zap.String("journey", "deleteUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		rest_err.NewBadRequestError("Invalid userId,must be a hex value")
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("journey", "deleteUser"),
		zap.String("userID", userId))

	c.Status(http.StatusOK)
}
