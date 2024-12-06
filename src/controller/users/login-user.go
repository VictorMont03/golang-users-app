package users

import (
	"net/http"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/validation"
	"github.com/VictorMont03/golang-users-app/src/controller/model/request"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "LoginUser"))
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding user request", err, zap.String("journey", "LoginUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(400, restErr)
		return
	}

	domain := model.NewUseLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUser(domain)

	if err != nil {
		logger.Error("Error trying to call LoginUser service", err, zap.String("journey", "LoginUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("End LoginUser controller", zap.String("journey", "LoginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
