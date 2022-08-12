package main

import (
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/infinitss13/InnoTaxiUser/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting application")
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error in loading env file : %v", err)
	}
	serverConfig := configs.NewServerConfig()

	port := serverConfig.SetTCPPort()
	server := new(entity.Server)
	err := server.Run(port, handler.SetRequestHandlers())
	if err != nil {
		logrus.Error(err)
		return
	}
}
