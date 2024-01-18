package service

import (
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	//ud é um objeto de service, um objeto vazio
	//userDomain é uma interface
	//mesmo conversando interfaces entre as nossas camadas, consegue pegar os valores lá de dentro.
	//não precisa deixar o model publico, podemos deixar controlavel. Disponibiliza os metodos que nós queremos
	//mas não disponibilza os objetos para ser alterado
	userDomainRepository, err := ud.userRepository.CreteUser(userDomain)
	if err != nil {
		logger.Error("Init trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
