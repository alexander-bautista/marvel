package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/alexander-bautista/marvel/pkg/character"
)

type characterMongoRepository struct {
	client  *mongo.Client
	timeout time.Duration
}

func NewMongoCharacterRepository(timeout int, client *mongo.Client) (character.CharacterRepository, error) {
	repo := &characterMongoRepository{
		client:  client,
		timeout: time.Duration(timeout) * time.Second,
	}

	return repo, nil
}

func (r *characterMongoRepository) GetOne(id int) (character *character.Character, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	collection := r.client.Database("todo").Collection("characters")

	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&character)

	if err != nil {
		return nil, errors.Wrap(err, "repository.Character.GetOne")
	}

	return character, err
}

func (r *characterMongoRepository) GetAll() ([]*character.Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	opts := options.Find()
	//opts.SetLimit(20)

	collection := r.client.Database("todo").Collection("characters")

	cursor, _ := collection.Find(ctx, bson.M{}, opts)

	defer cursor.Close(ctx)

	items := make([]*character.Character, 0)

	for cursor.Next(ctx) {
		oneItem := &character.Character{}
		err := cursor.Decode(&oneItem)

		if err != nil {
			return nil, err
		}

		items = append(items, oneItem)
	}

	return items, nil
}
