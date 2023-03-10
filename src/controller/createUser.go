package controller

import (
	"github.com/GabrielOGW/go-crud/src/configuration/logger"
	"github.com/GabrielOGW/go-crud/src/configuration/validation"
	"github.com/GabrielOGW/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("User Created succesfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
}
