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

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding user request", err, zap.String("journey", "CreateUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(400, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		logger.Error("Error creating user", err, zap.String("journey", "CreateUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("End CreateUser controller", zap.String("journey", "CreateUser"))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
