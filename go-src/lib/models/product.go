package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "products"

// This will be a product struct when we add such a collection.
// Test data only for now
type product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int32              `bson:"price,omitempty"`
	PriceID     string             `bson:"price_id,omitempty"`
	Stock       int32              `bson:"stock,omitempty"`
}

func GetProducts() ([]product, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	productCollection := GetMongoCollection(mongoClient, COLLECTION)
	var products []product

	selectOpts := options.Find()
	selectOpts.SetLimit(3)

	cursor, err := productCollection.Find(ctx, bson.M{}, selectOpts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}
