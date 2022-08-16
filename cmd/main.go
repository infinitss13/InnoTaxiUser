package main

import (
	"github.com/infinitss13/innotaxiuser"
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

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
