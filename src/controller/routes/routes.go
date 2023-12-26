package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/controller"
)

// InitRouters responsible to initialize the routes in main.go
// inicializar o router no main.go atrelar as rotas
func InitRouters(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
