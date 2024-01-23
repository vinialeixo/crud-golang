package repository

import (
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
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
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr

	FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
}
