package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbTimeout time.Duration = 3 * time.Second

func Init(env, url, database string) (*mongo.Database, error) {
	dbURL := url
	if env == "local" {
		dbURL += "&connect=direct"
	}
	clientOptions := options.Client().ApplyURI(dbURL)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(database)
	if err = createIndexes(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createIndexes(db *mongo.Database) error {
	if err := createBankIndexes(db); err != nil {
		return err
	}

	return nil
}

func createBankIndexes(db *mongo.Database) error {
	bankCollection := db.Collection("banks")
	codeNameIndex := mongo.IndexModel{
		Keys: bson.M{
			"code_name": 1, // index in ascending order
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := bankCollection.Indexes().CreateOne(ctx, codeNameIndex)
	if err != nil {
		return err
	}

	return nil
}
