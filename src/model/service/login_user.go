package service

import (
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {

	logger.Info("Init loginUser model", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))

	return user, token, nil
}
