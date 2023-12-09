package router

import (
	"github.com/Joybaruah/ssl-tracker/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func DomainRouter(domain *gin.RouterGroup) {

	// DOMAIN Controllers
	domain.GET("/domains", controllers.GetUsers)
	domain.POST("/domains", controllers.Login)

}
