package controller

import (
	"fmt"
	"log"

	"github.com/GabrielOGW/go-crud/src/configuration/validation"
	"github.com/GabrielOGW/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
			errRest := validation.ValidateUserError(err)

			c.JSON(errRest.Code, errRest)
			return
	}
	fmt.Println(userRequest)
}
