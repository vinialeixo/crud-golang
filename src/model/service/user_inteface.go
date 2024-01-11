package service

import (
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

// object
// dependencia?
type userDomainService struct {
	userRepository repository.UserRepository
}

// methods
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) //pq *UserDomain se caso der erro, não posso retornar objeto vazio, pq objeto vazio é diferente de nulo. Nulo quer dizer que dá erro
	DeleteUser(string) *rest_err.RestErr
}
