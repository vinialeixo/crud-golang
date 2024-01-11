package main

import (
	"github.com/vinialeixo/crud-golang/src/controller"
	"github.com/vinialeixo/crud-golang/src/model/repository"
	"github.com/vinialeixo/crud-golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	//Init dependices
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
