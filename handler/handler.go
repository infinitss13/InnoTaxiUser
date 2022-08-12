package handler

import (
	"github.com/infinitss13/InnoTaxiUser/cmd/logger"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/middleware"
)

type AuthHandlers struct {
	loggerMongo logger.LoggerMongo
	service     *services.Service
}

func NewAuthHandlers() *AuthHandlers {
	mongoDBClient, err := logger.NewClientMongo()
	if err != nil {
		return nil
	}
	srv := new(services.Service)
	srv.Db, err = dataBase.NewDB(configs.NewConfig())

	return &AuthHandlers{
		loggerMongo: logger.NewLogger(mongoDBClient),
		service:     srv,
	}
}

func SetRequestHandlers() *gin.Engine {
	router := gin.New()
	handler := NewAuthHandlers()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello message")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}
	api := router.Group("/api").Use(middleware.Auth())
	{
		api.GET("/rating", handler.getRating)
		api.POST("/order", orderTaxi)
		api.POST("/rateTrip", rateTrip)
	}
	return router
}
