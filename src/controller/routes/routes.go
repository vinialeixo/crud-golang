package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/controller"
	"github.com/vinialeixo/crud-golang/src/model"
)

// InitRouters responsible to initialize the routes in main.go
// inicializar o router no main.go atrelar as rotas
func InitRouters(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}
