package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/cmd/cache"
	"github.com/infinitss13/innotaxiuser/cmd/logger"
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/database"
	"github.com/infinitss13/innotaxiuser/middleware"
	"github.com/infinitss13/innotaxiuser/services"
)

type AuthHandlers struct {
	loggerMongo logger.LoggerMongo
	service     *services.Service
	cache       cache.RedisCash
}

func NewAuthHandlers() (*AuthHandlers, error) {
	mongoDBClient, err := logger.NewClientMongo()
	if err != nil {
		return nil, err
	}
	srv := new(services.Service)
	srv.Db, err = database.NewDB(configs.NewConfig())
	if err != nil {
		return nil, err
	}
	ch, err := cache.NewRedisCash()
	if err != nil {
		return nil, err
	}
	return &AuthHandlers{
		loggerMongo: logger.NewLogger(mongoDBClient),
		service:     srv,
		cache:       ch,
	}, nil
}
func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
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
	router.GET("/metrics", prometheusHandler())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}
	api := router.Group("/api").Use(middleware.Auth())
	{
		api.POST("/sign-out", handler.signOut)
		api.GET("/rating", handler.getRating)
		api.POST("/order", orderTaxi)
		api.POST("/rateTrip", rateTrip)
	}
	return router, nil
}
