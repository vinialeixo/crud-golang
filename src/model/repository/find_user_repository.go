package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/model"
	"github.com/vinialeixo/crud-golang/src/model/repository/entity"
	"github.com/vinialeixo/crud-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// FindUserByEmail e FindUserByID são metodos
func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dataBaseConnection.Collection(collection_name)

	//poder jogar o valor do banco de dados nesse objeto
	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity) //decode vai pegar o valor do banco e jogar dentro da variavel se não tiver nenhum problema

	//pode ter dois tipos de erro 1-não ter encontrado no bd / 2-erro user not found
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dataBaseConnection.Collection(collection_name)

	//poder jogar o valor do banco de dados nesse objeto
	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity) //decode vai pegar o valor do banco e jogar dentro da variavel se não tiver nenhum problema

	//pode ter dois tipos de erro 1-não ter encontrado no bd / 2-erro user not found
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
