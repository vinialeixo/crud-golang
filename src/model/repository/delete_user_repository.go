package repository

import (
	"context"
	"os"

	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Init CreteUser repository",
		zap.String("journey", "CreteUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dataBaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "createUser"))
		return nil
	}

	logger.Info(
		"updateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
