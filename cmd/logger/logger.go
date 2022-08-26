package logger

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/configs"
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

func (d LoggerMongo) LogError(ctx *gin.Context, err error) error {
	timeNow := time.Now()
	timeNow.Format(time.RFC3339)
	doc := bson.D{
		primitive.E{
			Key: "loglevel", Value: "ERROR",
		},
		primitive.E{
			Key:   "requestType",
			Value: ctx.Request.Method,
		},
		primitive.E{
			Key:   "error",
			Value: err.Error(),
		},
		primitive.E{
			Key: "requestTime", Value: timeNow.Format(time.RFC3339),
		},
	}
	//doc := bson.D{
	//
	//	{
	//		"logLevel", "ERROR",
	//	},
	//	{
	//		"requestType", ctx.Request.Method,
	//	},
	//	{
	//		"error", err.Error(),
	//	},
	//	{
	//		"requestTime", timeNow.Format(time.RFC3339),
	//	},
	//}
	_, err = d.collection.InsertOne(context.TODO(), doc)
	if err != nil {

		return err
	}
	return nil
}
func (d LoggerMongo) LogInfo(ctx *gin.Context) error {
	timeNow := time.Now()
	timeNow.Format(time.RFC3339)
	doc := bson.D{
		{
			"logLevel", "INFO",
		},
		{
			"requestType", ctx.Request.Method,
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
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(newConnection.MongoDBName), nil
}
