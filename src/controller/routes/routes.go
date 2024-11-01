package routes

import (
	"github.com/VictorMont03/golang-users-app/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, uc users.UserControllerInterface) {
	r.GET("/users/:id", uc.GetUserById)
	// r.GET("/users", users.QueryUsers, )
	r.POST("/users", uc.CreateUser)
	r.PUT("/users/:id", uc.UpdateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
}
