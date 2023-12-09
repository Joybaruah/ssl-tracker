package router

import (
	"github.com/Joybaruah/ssl-tracker/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(user *gin.RouterGroup) {

	// USER Controllers
	user.GET("/users", controllers.GetUsers)
	user.POST("/login", controllers.Login)
	user.POST("/register", controllers.Register)

}
