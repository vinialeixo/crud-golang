package repository

import (
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(databse *mongo.Database) UserRepository {
	return &userRepository{
		databse,
	}
}

// userRepository é o objeto
type userRepository struct {
	dataBaseConnection *mongo.Database
}
type UserRepository interface {
	CreteUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
