package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/crud-golang/src/controller/dto"
	"github.com/vinialeixo/crud-golang/src/controller/validation"
)

func CreateUser(c *gin.Context) {

	var userRequest dto.UserRequest

	//&userRequest To be able to modify the original userRequest variable (defined outside the function) with the data from the request,
	//you need to pass a pointer to that variable
	//By using &userRequest, you are passing the memory address of the userRequest variable to the ShouldBind function. This allows ShouldBind to modify the actual variable's content,
	err := c.ShouldBind(&userRequest)
	if err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error()) //string message
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(userRequest)

}
