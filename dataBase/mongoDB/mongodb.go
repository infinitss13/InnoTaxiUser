package mongoDB

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type ConnectionMongo struct {
	host     string
	port     string
	username string
	password string
	dbname   string
	authdb   string
}

func NewConnectionMongo() ConnectionMongo {
	return ConnectionMongo{
		host:     os.Getenv("HOST_MONGO"),
		port:     os.Getenv("PORT_MONGO"),
		username: os.Getenv("USERNAME_MONGO"),
		password: os.Getenv("PASSWORD_MONGO"),
		dbname:   os.Getenv("DBNAME_MONGO"),
		authdb:   os.Getenv("AUTH_MONGO"),
	}
}

func NewClientMongo(context *gin.Context) (db *mongo.Database, err error) {
	var mongoDBURL string
	var isAuth bool
	newConnection := NewConnectionMongo()
	if newConnection.username == "" && newConnection.password == "" {

		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", newConnection.host, newConnection.port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			newConnection.username, newConnection.password, newConnection.host, newConnection.port)
	}
	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		if newConnection.authdb == "" {
			newConnection.authdb = newConnection.dbname
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: newConnection.authdb,
			Username:   newConnection.username,
			Password:   newConnection.password,
		})
	}
	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB : %v", err)
	}
	err = client.Ping(context, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB : %v", err)
	}

	return client.Database(newConnection.dbname), nil
}
