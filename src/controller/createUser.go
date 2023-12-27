package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/controller/dto"
)

func CreateUser(c *gin.Context) {

	var userRequest dto.UserRequest

	//&userRequest vai jogar dentro do endereço de memoria
	err := c.ShouldBind(&userRequest)
	if err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error())) //string message
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
