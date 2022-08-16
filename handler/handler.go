package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/cmd/logger"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/database"
	"github.com/infinitss13/InnoTaxiUser/middleware"
	"github.com/infinitss13/InnoTaxiUser/services"
)

type AuthHandlers struct {
	loggerMongo logger.LoggerMongo
	service     *services.Service
}

func NewAuthHandlers() (*AuthHandlers, error) {
	mongoDBClient, err := logger.NewClientMongo()
	if err != nil {
		return nil, err
	}
	srv := new(services.Service)
	srv.Db, err = database.NewDB(configs.NewConfig())

	return &AuthHandlers{
		loggerMongo: logger.NewLogger(mongoDBClient),
		service:     srv,
	}, nil
}

func SetRequestHandlers() (*gin.Engine, error) {
	router := gin.New()
	handler, err := NewAuthHandlers()
	if err != nil {
		return nil, err
	}
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
	return router, nil
}
