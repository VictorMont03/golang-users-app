package users

import (
	"net/http"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/controller/model/response"
	"github.com/VictorMont03/golang-users-app/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) QueryUsers(c *gin.Context) {
	logger.Info("Init QueryUsers Controller", zap.String("journey", "QueryUsers"))

	params := c.Request.URL.Query()

	userDomain, _ := uc.service.QueryUser(params)

	var responseDomain []response.UserResponse

	for _, user := range userDomain {
		responseDomain = append(responseDomain, view.ConvertDomainToResponse(user))
	}

	logger.Info("End QueryUsers Controller", zap.String("journey", "QueryUsers"))
	c.JSON(http.StatusOK, responseDomain)
}
