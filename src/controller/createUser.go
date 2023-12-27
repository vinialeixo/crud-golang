package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
	"github.com/vinialeixo/crud-golang/src/controller/dto"
)

func CreateUser(c *gin.Context) {

	var userRequest dto.UserRequest

	//&userRequest To be able to modify the original userRequest variable (defined outside the function) with the data from the request,
	//you need to pass a pointer to that variable
	//By using &userRequest, you are passing the memory address of the userRequest variable to the ShouldBind function. This allows ShouldBind to modify the actual variable's content,
	err := c.ShouldBind(&userRequest)
	if err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error())) //string message
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

}
