package redisdb

import (
	"context"
	"os"

	"github.com/piqba/mtss-go/pkg/errors"
	"github.com/piqba/mtss-go/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrMongoDbConnection = errors.NewError("Mongo: Fail to connection")
	ErrMongoDbCheckConn  = errors.NewError("Mongo: Fail to check connection")
)

func GetMongoDbClient() *mongo.Client {
	var clientInstance *mongo.Client

	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.LogError(ErrMongoDbConnection.Error())

	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.LogError(ErrMongoDbCheckConn.Error())
	}
	clientInstance = client
	return clientInstance
}
