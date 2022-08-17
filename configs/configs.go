package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

type ConnectionRedis struct {
	RedisHost    string
	RedisDB      int
	RedisExpires time.Duration
}

type ServerConfig struct {
	tcpPort string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		tcpPort: os.Getenv("PORT"),
	}
}
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func NewConfig() DBConfig {
	return DBConfig{
		DBHost:     GetEnv("HOST_DB", "localhost"),
		DBPort:     GetEnv("PORT_DB", "5436"),
		DBUsername: GetEnv("USERNAME_DB", "postgres"),
		DBName:     GetEnv("DBNAME_DB", "postgres"),
		DBSslmode:  GetEnv("SSLMODE_DB", "disable"),
		DBPassword: GetEnv("PASSWORD_DB", "qwerty"),
	}
}

func NewConnectionMongo() ConnectionMongo {
	return ConnectionMongo{
		MongoHost:       GetEnv("HOST_MONGO", "127.0.0.1"),
		MongoPort:       GetEnv("PORT_MONGO", "27017"),
		MongoDBName:     GetEnv("DBNAME_MONGO", "projectdb"),
		MongoCollection: GetEnv("COLLECTION_MONGO", "logging"),
	}
}

func NewConnectionRedis() (ConnectionRedis, error) {
	DBNumber, err := strconv.Atoi(GetEnv("DB_REDIS", "0"))
	if err != nil {
		return ConnectionRedis{}, err
	}
	timeInt, err := strconv.Atoi(GetEnv("TOKEN_EXPIRES", "15"))
	if err != nil {
		return ConnectionRedis{}, err
	}
	timeExpires := time.Duration(timeInt) * time.Minute
	//TODO : realize getting time from env file
	return ConnectionRedis{
		RedisHost:    GetEnv("HOST_REDIS", "127.0.0.1:6379"),
		RedisDB:      DBNumber,
		RedisExpires: timeExpires,
	}, nil
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
