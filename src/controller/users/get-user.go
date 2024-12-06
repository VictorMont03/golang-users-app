package users

import (
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) GetUserById(c *gin.Context) {
	logger.Info("Init GetUserById Controller", zap.String("journey", "controller GetUserById"))

	userId := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error inside GetUserById Controller / Invalid ID", err, zap.String("journey", "controller GetUserById"))
		c.JSON(400, rest_err.NewBadRequestError("Invalid ID"))
		return
	}

	userDomain, err := uc.service.GetUser(userId)

	if err != nil {
		logger.Error("Error inside GetUserById Controller", err, zap.String("journey", "controller GetUserById"))
		c.JSON(err.Code, err)
	}

	logger.Info("End GetUserById Controller", zap.String("journey", "controller GetUserById"), zap.String("user_id", userDomain.GetID()))

	c.JSON(200, view.ConvertDomainToResponse(userDomain))
}
