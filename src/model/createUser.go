package model

import (
	"github.com/GabrielOGW/go-crud/src/configuration/logger"
	"github.com/GabrielOGW/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}


