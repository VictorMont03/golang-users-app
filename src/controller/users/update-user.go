package users

import (
	"net/http"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/validation"
	"github.com/VictorMont03/golang-users-app/src/controller/model/request"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "UpdateUser"))
	var userRequest request.UserUpdateRequest

	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding user request", err, zap.String("journey", "UpdateUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(400, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)

	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err, zap.String("journey", "UpdateUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("End UpdateUser controller", zap.String("journey", "UpdateUser"), zap.String("userId", userId))

	c.Status(http.StatusOK)
}
