package controller

import (
	"github.com/GabrielOGW/go-crud/src/configuration/logger"
	"github.com/GabrielOGW/go-crud/src/configuration/validation"
	"github.com/GabrielOGW/go-crud/src/controller/model/request"
	"github.com/GabrielOGW/go-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(userRequest.Email, 
		userRequest.Password, 
		userRequest.Name, 
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User Created succesfully",
		zap.String("journey", "createUser"),

		c.JSON(http.StatusOK, "")
		)
}
