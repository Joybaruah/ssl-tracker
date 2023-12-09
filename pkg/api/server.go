package api

import (
	middlewares "github.com/Joybaruah/ssl-tracker/pkg/middleware"
	"github.com/Joybaruah/ssl-tracker/pkg/router"
	"github.com/gin-gonic/gin"
)

func APIService() {

	route := gin.New()

	route.Use(gin.Logger())
	// route.Use(middlewares.RequireApiAuth)

	// API Group
	api := route.Group("/api")
	api.Use(middlewares.RequireApiAuth)

	// Route Groups
	user := api.Group("/user")
	domain := api.Group("/domain")

	// Route Handlers
	router.UserRouter(user)
	router.DomainRouter(domain)

	route.Run()

}
