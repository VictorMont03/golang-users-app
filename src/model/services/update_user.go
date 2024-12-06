package services

import (
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser Model", zap.String("journey", "model UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error inside UpdateUser function Model/Service", err, zap.String("journey", "model UpdateUser"))
		return err
	}

	logger.Info("End UpdateUser Model/Service", zap.String("journey", "model UpdateUser"), zap.String("user_id", userId))

	return nil
}
