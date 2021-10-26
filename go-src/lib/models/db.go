package models

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/keyprez/keyprez/go-src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	atlasUri := utils.GetEnvVar("ATLAS_URI")
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		client, err := mongo.NewClient(options.Client().ApplyURI(atlasUri))
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func GetMongoCollection(collection string) *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	db := utils.GetEnvVar("ATLAS_DB")
	return client.Database(db).Collection(collection)
}
