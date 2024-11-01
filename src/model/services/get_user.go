package services

import (
	"fmt"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) GetUser(string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetUser Model", zap.String("journey", "model GetUser"))

	fmt.Println(ud)

	return nil, nil
}
