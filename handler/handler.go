package handler

import (
	"net/http"

	"github.com/infinitss13/InnoTaxiUser/middleware"

	"github.com/gin-gonic/gin"
)

func SetRequestHandlers() *gin.Engine {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello message")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signUp)
		auth.POST("/sign-in", signIn)
	}
	api := router.Group("/api").Use(middleware.Auth())
	{
		api.GET("/rating", getRating)
		api.POST("/order", orderTaxi)
		api.POST("/rateTrip", rateTrip)
	}
	return router
}
