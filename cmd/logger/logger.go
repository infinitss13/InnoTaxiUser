package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type Logger interface {
	LogError(ctx *gin.Context, err error) error
	LogInfo(ctx *gin.Context) error
}

func NewClientMongo(ctx context.Context) (db *mongo.Database, err error) {
	var mongoDBURL string
	newConnection := configs.NewConnectionMongo()

	mongoDBURL = fmt.Sprintf("mongodb://%s:%s", newConnection.MongoHost, newConnection.MongoPort)
	clientOptions := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(newConnection.MongoDBName), nil
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
	_, err = d.collection.InsertOne(context.TODO(), doc)
	return err
}
func (d LoggerMongo) LogInfo(ctx *gin.Context) error {
	timeNow := time.Now()
	timeNow.Format(time.RFC3339)
	doc := bson.D{
		primitive.E{
			Key: "loglevel", Value: "INFO",
		},
		primitive.E{
			Key:   "requestType",
			Value: ctx.Request.Method,
		},
		primitive.E{
			Key:   "error",
			Value: "no",
		},
		primitive.E{
			Key: "requestTime", Value: timeNow.Format(time.RFC3339),
		},
	}

	_, err := d.collection.InsertOne(context.TODO(), doc)
	return err
}
