package configs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	hostDB     string
	portDB     string
	usernameDB string
	nameDB     string
	sslModeDB  string
	passwordDB string
}

type ServerConfig struct {
	tcpPort string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		tcpPort: os.Getenv("PORT"),
	}
}

func NewConfig() *DBConfig {
	return &DBConfig{

		hostDB:     os.Getenv("HOST_DB"),
		portDB:     os.Getenv("PORT_DB"),
		usernameDB: os.Getenv("USERNAME_DB"),
		nameDB:     os.Getenv("DBNAME_DB"),
		sslModeDB:  os.Getenv("SSLMODE_DB"),
		passwordDB: os.Getenv("PASSWORD_DB"),
	}
}

func (c *DBConfig) ConnectionDbData() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.hostDB, c.portDB, c.usernameDB, c.nameDB, c.passwordDB, c.sslModeDB)
}

func (c *ServerConfig) SetTCPPort() string {
	var port string
	port = ":" + c.tcpPort
	if port == ":" {
		port = ":8080"
		logrus.Info("Switching to default port: 8080")
	}
	return port
}
