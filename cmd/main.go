package main

import (
	"context"
	"github.com/infinitss13/innotaxiuser"
	"github.com/infinitss13/innotaxiuser/cmd/cache"
	"github.com/infinitss13/innotaxiuser/cmd/logger"
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/database"
	_ "github.com/infinitss13/innotaxiuser/docs"
	"github.com/infinitss13/innotaxiuser/handler"
	"github.com/infinitss13/innotaxiuser/services"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title InnoTaxi user service
// @version 1.0
// @description API server for InnoTaxi user service

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.Info("Starting  XUINYA VASH DOCKER application")
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error in loading env file : %v", err)
	}
	serverConfig := configs.NewServerConfig()

	port := serverConfig.SetTCPPort()
	server := new(innotaxiuser.Server)

	ctx := context.Background()

	mongoDBClient, err := logger.NewClientMongo(ctx)
	if err != nil {
		logrus.Errorf("error in connection mongo : %v", err)
	}
	log := logger.NewLogger(mongoDBClient)

	db, err := database.NewDataBase(configs.NewDBConfig())
	if err != nil {
		logrus.Errorf("error in connection postgres : %v", err)
	}
	srv := services.NewService(db)

	cacheRedis, err := cache.NewRedisCache()
	if err != nil {
		logrus.Errorf("error in connection cache : %v", err)
	}
	handlers, err := handler.SetRequestHandlers(log, srv, cacheRedis)
	if err != nil {
		logrus.Errorf("error setting http handlers %v", err)
		return
	}
	err = server.Run(port, handlers)
	if err != nil {
		logrus.Error(err)
		return
	}

}
