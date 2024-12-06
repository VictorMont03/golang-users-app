package services

import (
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) GetUser(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetUser Model", zap.String("journey", "model GetUser"))

	return ud.userRepository.GetUser(id)
}
