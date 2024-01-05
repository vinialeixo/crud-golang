package service

import (
	"fmt"

	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	//ud é um objeto de service, um objeto vazio
	//userDomain é uma interface
	//mesmo conversando interfaces entre as nossas camadas, consegue pegar os valores lá de dentro.
	//não precisa deixar o model publico, podemos deixar controlavel. Disponibiliza os metodos que nós queremos
	//mas não disponibilza os objetos para ser alterado
	fmt.Println(userDomain.GetPassword())
	return nil
}
