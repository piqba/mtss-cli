package cli

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piqba/mtss-cli/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrMongoDbConnection = internal.NewError("Mongo: Fail to connection")
	ErrMongoDbCheckConn  = internal.NewError("Mongo: Fail to check connection")
	ErrRedisDbCheckConn  = internal.NewError("Redis: Fail to check connection")
)

func GetMongoDbClient() *mongo.Client {
	var clientInstance *mongo.Client

	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		internal.LogError(ErrMongoDbConnection.Error())

	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		internal.LogError(ErrMongoDbCheckConn.Error())
	}
	clientInstance = client
	return clientInstance
}

func GetRedisDbClient() *redis.Client {

	clientInstance := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URI"),  // use default Addr
		Password:     os.Getenv("REDIS_PASS"), // no password set
		DB:           0,
		DialTimeout:  60 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	})

	_, err := clientInstance.Ping(context.TODO()).Result()
	if err != nil {
		internal.LogError(ErrRedisDbCheckConn.Error())
	}
	return clientInstance
}
