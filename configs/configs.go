package configs

import (
	"fmt"
	"github.com/joho/godotenv"
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

func NewConfig() DBConfig {
	return DBConfig{
		DBHost:     os.Getenv("HOST_DB"),
		DBPort:     os.Getenv("PORT_DB"),
		DBUsername: os.Getenv("USERNAME_DB"),
		DBName:     os.Getenv("DBNAME_DB"),
		DBSslmode:  os.Getenv("SSLMODE_DB"),
		DBPassword: os.Getenv("PASSWORD_DB"),
	}

}
func NewConnectionMongo() ConnectionMongo {
	if err := godotenv.Load(); err != nil {
		return ConnectionMongo{
			MongoHost:       "127.0.0.1",
			MongoPort:       "27017",
			MongoDBName:     "projectdb",
			MongoCollection: "logging",
		}
	} else {
		return ConnectionMongo{
			MongoHost:       os.Getenv("HOST_MONGO"),
			MongoPort:       os.Getenv("PORT_MONGO"),
			MongoDBName:     os.Getenv("DBNAME_MONGO"),
			MongoCollection: os.Getenv("COLLECTION_MONGO"),
		}
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
