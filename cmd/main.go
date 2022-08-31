package main

import (
	"github.com/infinitss13/innotaxiuser"
	"github.com/infinitss13/innotaxiuser/configs"
	_ "github.com/infinitss13/innotaxiuser/docs"
	"github.com/infinitss13/innotaxiuser/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title InnoTaxi user service
// @version 1.0
// @description API server for InnoTaxi user service

// @host localhost:8000
// @BasePath /

// @securityDefinition ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.Info("Starting application")
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error in loading env file : %v", err)
	}
	serverConfig := configs.NewServerConfig()

	port := serverConfig.SetTCPPort()
	server := new(innotaxiuser.Server)
	handlers, err := handler.SetRequestHandlers()
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
