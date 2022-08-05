package main

import (
	"fmt"
	"log"
	"os"

	"github.com/infinitss13/InnoTaxiUser"
	"github.com/infinitss13/InnoTaxiUser/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting application")
	if err := godotenv.Load(); err != nil {
		log.Fatal("error in loading env file")
	}
	port := ":" + os.Getenv("PORT")

	if port == "" {
		port = ":8080"
		log.Println("Switching to default port: 8080")
	}
	fmt.Println(port)
	handlers := handler.NewHandler()
	server := new(InnoTaxiUser.Server)
	err := server.Run(port, handlers.SetRequestHandlers())
	if err != nil {
		logrus.Error(err)
	}

}
