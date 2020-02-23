package mongo

import (
	"context"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() (*mongo.Client, error) {

	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"
	}
	mongodb := os.Getenv("MONGO_DB")
	if mongodb == "" {
		mongodb = "todo"
	}

	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

	if mongoTimeout == 0 {
		mongoTimeout = 10
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
