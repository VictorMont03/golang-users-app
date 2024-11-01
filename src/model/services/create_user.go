package services

import (
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser Model", zap.String("journey", "model CreateUser"))

	userDomain.EncryptPassword()

	urr, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error creating user", err)
		return nil, err
	}

	return urr, nil
}
