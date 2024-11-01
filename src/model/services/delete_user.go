package services

import (
	"fmt"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init DeleteUser Model", zap.String("journey", "model DeleteUser"))

	fmt.Println(ud)

	return nil
}
