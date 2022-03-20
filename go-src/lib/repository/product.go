package repository

import (
	"context"
	"time"

	"github.com/keyprez/keyprez/go-src/lib/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProduct() ([]models.Product, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	productCollection := GetMongoCollection(mongoClient, PRODUCT_COLLECTION)
	var products []models.Product

	selectOpts := options.Find()
	selectOpts.SetLimit(1)

	cursor, err := productCollection.Find(ctx, bson.M{}, selectOpts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProducts() ([]models.Product, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	productCollection := GetMongoCollection(mongoClient, PRODUCT_COLLECTION)
	var products []models.Product

	selectOpts := options.Find()
	selectOpts.SetLimit(4)

	cursor, err := productCollection.Find(ctx, bson.M{}, selectOpts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProductStock(priceId string, quantity int64) (*mongo.UpdateResult, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, PRODUCT_COLLECTION)

	res, updateErr := col.UpdateOne(ctx, bson.M{"price_id": priceId}, bson.D{{"$inc", bson.D{{"Stock", quantity}}}})
	if updateErr != nil {
		return nil, updateErr
	}

	return res, nil

}
