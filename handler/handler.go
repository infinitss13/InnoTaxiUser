package handler

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/cmd/cache"
	"github.com/infinitss13/innotaxiuser/cmd/logger"
	"github.com/infinitss13/innotaxiuser/services"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
)

type AuthHandlers struct {
	LoggerMongo logger.Logger
	UserService services.UserService
	Cache       cache.Cache
}

//func NewAuthHandlers(log logger.LoggerMongo, srv services.Service, cache cache.RedisCash) *AuthHandlers {
//	return &AuthHandlers{
//		LoggerMongo: log,
//		UserService: srv,
//		Cache:       cache,
//	}
//}
func NewAuthHandlers(log logger.Logger, srv services.UserService, cache cache.Cache) *AuthHandlers {
	return &AuthHandlers{
		LoggerMongo: log,
		UserService: srv,
		Cache:       cache,
	}
}

func SetRequestHandlers(log logger.Logger, srv services.UserService, cache cache.Cache) (*gin.Engine, error) {
	router := gin.New()
	handler := NewAuthHandlers(log, srv, cache)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello message")
	})
	pprof.Register(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", prometheusHandler())
	auth := router.Group("/auth").Use(metricHttpStatus)
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api").Use(metricHttpStatus).Use(handler.UserService.Auth())
	{

		api.GET("/profile", handler.getProfile)
		api.PATCH("/profile", handler.updateProfile)
		api.DELETE("/profile", handler.deleteProfile)
		api.GET("/sign-out", handler.signOut)
		api.GET("/rating", handler.getRating)
		api.POST("/order", orderTaxi)
		api.POST("/rateTrip", rateTrip)
	}
	return router, nil
}

func (handler AuthHandlers) GetAndCheckToken(context *gin.Context) (string, error) {
	tokenSigned, err := handler.UserService.GetToken(context)
	if err != nil {
		errorToken := fmt.Errorf("get rating error : %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger", errorLogger.Error())
		}
		HandleError(err, context)
		return "", errorToken
	}
	isKey, err := handler.Cache.GetValue(tokenSigned)
	if isKey && err != cache.UserSignedOut {
		context.JSON(http.StatusBadRequest, "user with this token signed-out")
		return "", err
	}
	return tokenSigned, nil
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
