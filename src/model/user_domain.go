package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
)

// NewUserDomain constructor
func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil)) //não precisa retornar nada, pq vai alterar no ponteiro do objeto
}

// methods
type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr) //pq *UserDomain se caso der erro, não posso retornar objeto vazio, pq objeto vazio é diferente de nulo. Nulo quer dizer que dá erro
	DeleteUser(string) *rest_err.RestErr
}
