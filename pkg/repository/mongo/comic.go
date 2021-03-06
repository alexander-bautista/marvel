package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/alexander-bautista/marvel/pkg/comic"
)

type comicMongoRepository struct {
	client  *mongo.Client
	timeout time.Duration
}

func NewMongoComicRepository(timeout int, client *mongo.Client) (comic.ComicRepository, error) {
	repo := &comicMongoRepository{
		client:  client,
		timeout: time.Duration(timeout) * time.Second,
	}

	return repo, nil
}

func (r *comicMongoRepository) GetOne(id int) (comic *comic.Comic, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	collection := r.client.Database("todo").Collection("comics")

	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, errors.Wrap(err, "repository.Comic.GetOne")
	}

	return comic, err
}

func (r *comicMongoRepository) GetAll() ([]*comic.Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	opts := options.Find()
	//opts.SetLimit(20)

	collection := r.client.Database("todo").Collection("comics")

	cursor, _ := collection.Find(ctx, bson.M{}, opts)

	defer cursor.Close(ctx)

	items := make([]*comic.Comic, 0)

	for cursor.Next(ctx) {
		oneItem := &comic.Comic{}
		err := cursor.Decode(&oneItem)

		if err != nil {
			return nil, err
		}

		items = append(items, oneItem)
	}

	return items, nil
}
