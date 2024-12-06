package services

import (
	"net/url"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) QueryUser(params url.Values) ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init QueryUser Model", zap.String("journey", "model QueryUser"))

	return ud.userRepository.QueryUser(params)
}
