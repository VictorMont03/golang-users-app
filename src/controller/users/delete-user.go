package users

import (
	"net/http"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", "DeleteUser"))
	userId := c.Param("id")

	err := uc.service.DeleteUser(userId)

	if err != nil {
		logger.Error("Error trying to call DeleteUser service", err, zap.String("journey", "DeleteUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("End DeleteUser controller", zap.String("journey", "DeleteUser"), zap.String("userId", userId))

	c.Status(http.StatusOK)
}
