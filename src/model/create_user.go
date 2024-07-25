package model

import (
	"fmt"
	"teste/src/configuration/logger"
	"teste/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Iniciando Criação", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
