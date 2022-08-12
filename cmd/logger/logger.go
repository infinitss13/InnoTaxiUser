package logger

import (
	"context"
	"fmt"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoggerMongo struct {
	collection *mongo.Collection
}

func NewLogger(database *mongo.Database) LoggerMongo {
	return LoggerMongo{
		collection: database.Collection(configs.NewConnectionMongo().MongoCollection),
	}
}

func (d LoggerMongo) AddNewErrorLog(ctx *gin.Context, userPhone string, errorRequest error, comments string) error {
	timeNow := time.Now()
	timeNow.Format(time.RFC3339)
	doc := bson.D{
		{
			"logLevel", "ERROR",
		},
		{
			"message", comments,
		},
		{
			"requestType", ctx.Request.Method,
		},
		{
			"userPhone", userPhone,
		},

		{
			"error", errorRequest.Error(),
		},
		{
			"requestTime", timeNow.Format(time.RFC3339),
		},
	}

	_, err := d.collection.InsertOne(context.TODO(), doc)
	if err != nil {

		return err
	}
	return nil
}

func (d LoggerMongo) AddNewInfoLog(ctx *gin.Context, userPhone string, comments string) error {
	timeNow := time.Now()
	timeNow.Format(time.RFC3339)
	doc := bson.D{
		{
			"logLevel", "INFO",
		},
		{
			"message", comments,
		},
		{
			"requestType", ctx.Request.Method,
		},
		{
			"userPhone", userPhone,
		},

		{
			"error", "no",
		},
		{
			"requestTime", timeNow.Format(time.RFC3339),
		},
	}

	_, err := d.collection.InsertOne(context.TODO(), doc)
	if err != nil {

		return err
	}
	return nil
}

func NewClientMongo() (db *mongo.Database, err error) {
	var mongoDBURL string
	newConnection := configs.NewConnectionMongo()

	mongoDBURL = fmt.Sprintf("mongodb://%s:%s", newConnection.MongoHost, newConnection.MongoPort)

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB : %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB : %v", err)
	}

	return client.Database(newConnection.MongoDBName), nil
}
