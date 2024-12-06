package users

import (
	"github.com/VictorMont03/golang-users-app/src/model/services"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface services.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	QueryUsers(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
