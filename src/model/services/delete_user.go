package services

import (
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser Model", zap.String("journey", "model DeleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error("Error inside DeleteUser function Model/Service", err, zap.String("journey", "model DeleteUser"))
		return err
	}

	logger.Info("End DeleteUser Model/Service", zap.String("journey", "model DeleteUser"), zap.String("user_id", userId))

	return nil
}
