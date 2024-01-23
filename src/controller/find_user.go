package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserById controller", zap.String("journey", "findUserById"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
	}
	logger.Info(fmt.Sprintf("User authenticated: %v", user))

	userId := c.Param("userId")
	fmt.Println(userId)
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validated userId", err, zap.String("journey", "findUserById"))
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserBy Id services", err, zap.String("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed sucessfully", zap.String("journey", "findUserById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
	}
	logger.Info(fmt.Sprintf("User authenticated: %v", user))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validated userId", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserBy Email services", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed sucessfully", zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
