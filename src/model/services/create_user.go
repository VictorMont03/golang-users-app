package services

import (
	"net/url"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser Model/Service", zap.String("journey", "model CreateUser"))

	params := url.Values{}
	params.Add("email", userDomain.GetEmail())

	preUsers, err := ud.userRepository.QueryUser(params)

	if err != nil {
		logger.Error("Error inside createUser function Model/Service, Search for previous email failed", err, zap.String("journey", "model CreateUser"))
		return nil, err
	}

	if len(preUsers) != 0 {
		logger.Error("Error inside createUser function Model/Service, Email already exists", nil, zap.String("journey", "model CreateUser"))
		return nil, rest_err.NewBadRequestError("Email already exists")
	}

	userDomain.EncryptPassword()

	urr, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error inside createUser function Model/Service", err, zap.String("journey", "model CreateUser"))
		return nil, err
	}

	logger.Info("End CreateUser Model/Service", zap.String("journey", "model CreateUser"), zap.String("user_id", urr.GetID()))

	return urr, nil
}
