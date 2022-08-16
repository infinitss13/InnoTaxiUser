package configs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBName     string
	DBSslmode  string
	DBPassword string
}

type ConnectionMongo struct {
	MongoHost       string
	MongoPort       string
	MongoDBName     string
	MongoCollection string
}

type ServerConfig struct {
	tcpPort string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		tcpPort: os.Getenv("PORT"),
	}
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func NewConfig() DBConfig {
	return DBConfig{
		DBHost:     getEnv("HOST_DB", "localhost"),
		DBPort:     getEnv("PORT_DB", "5436"),
		DBUsername: getEnv("USERNAME_DB", "postgres"),
		DBName:     getEnv("DBNAME_DB", "postgres"),
		DBSslmode:  getEnv("SSLMODE_DB", "disable"),
		DBPassword: getEnv("PASSWORD_DB", "qwerty"),
	}
}

func NewConnectionMongo() ConnectionMongo {
	return ConnectionMongo{
		MongoHost:       getEnv("HOST_MONGO", "127.0.0.1"),
		MongoPort:       getEnv("PORT_MONGO", "27017"),
		MongoDBName:     getEnv("DBNAME_MONGO", "projectdb"),
		MongoCollection: getEnv("COLLECTION_MONGO", "logging"),
	}
}

func (c *DBConfig) ConnectionDbData() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUsername, c.DBName, c.DBPassword, c.DBSslmode)
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
