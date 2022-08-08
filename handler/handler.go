package handler

import (
	"github.com/infinitss13/InnoTaxiUser/middleware"
	"net/http"

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
		api.GET("/ping", GetRating)
	}
	return router
}
