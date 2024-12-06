package services

import (
	"net/url"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser Model/Service", zap.String("journey", "model LoginUser"))

	userDomain.EncryptPassword()

	params := url.Values{}
	params.Add("email", userDomain.GetEmail())
	params.Add("password", userDomain.GetPassword())

	logger.Info("Params", zap.String("params", params.Encode()))

	userArr, err := ud.userRepository.QueryUser(params)

	if err != nil {
		logger.Error("Error inside LoginUser function Model/Service, Search for user failed", err, zap.String("journey", "model LoginUser"))
		return nil, "", err
	}

	if len(userArr) == 0 {
		logger.Error("Error inside LoginUser function Model/Service, Email not found", nil, zap.String("journey", "model LoginUser"))
		return nil, "", rest_err.NewBadRequestError("Incorrect password or email")
	}

	user := userArr[0]

	token, err := user.GenerateToken()

	logger.Info("End LoginUser Model/Service", zap.String("journey", "model LoginUser"), zap.String("user_id", user.GetID()))

	return user, token, nil
}
