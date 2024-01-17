package service

import (
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	//conectar service no repository
	logger.Info("Init FindUserByEmail services", zap.String("journey", "FindUserByEmail"))

	userEmail, err := ud.userRepository.FindUserByEmail(email)
	//conectar o controler no service
	if err != nil {
		logger.Error("Init FindUserByEmail model", err, zap.String("journey", "FindUserByEmail"))
		return nil, err
	}

	return userEmail, nil
}

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	//conectar service no repository
	logger.Info("Init findUserByID services", zap.String("journey", "findUserByID"))

	userID, err := ud.userRepository.FindUserByID(id)
	//conectar o controler no service
	if err != nil {
		logger.Error("Init findUserByID services", err, zap.String("journey", "findUserByID"))
		return nil, err
	}

	return userID, nil
}
