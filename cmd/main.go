package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinialeixo/crud-golang/src/configuration/db/mongodb"
	"github.com/vinialeixo/crud-golang/src/configuration/logger"
	"github.com/vinialeixo/crud-golang/src/controller"
	"github.com/vinialeixo/crud-golang/src/controller/routes"
	"github.com/vinialeixo/crud-golang/src/model/repository"
	"github.com/vinialeixo/crud-golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect do database, error = %s \n", err.Error())
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRouters(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	//Init dependices
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
