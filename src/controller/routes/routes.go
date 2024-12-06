package routes

import (
	"github.com/VictorMont03/golang-users-app/src/controller/users"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, uc users.UserControllerInterface) {
	r.GET("/users/:id", uc.GetUserById)
	r.GET("/users", model.AuthorizationMiddleware, uc.QueryUsers)
	r.POST("/users", uc.CreateUser)
	r.PUT("/users/:id", uc.UpdateUser)
	r.DELETE("/users/:id", uc.DeleteUser)

	// Login
	r.POST("/login", uc.LoginUser)
}
