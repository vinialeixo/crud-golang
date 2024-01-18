package service

import (
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Init trying to call repository", err, zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
