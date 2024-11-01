package services

import (
	"fmt"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser Model", zap.String("journey", "model UpdateUser"))

	fmt.Println(ud)

	return nil
}
