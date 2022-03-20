package repository

import (
	"context"
	"time"

	"github.com/keyprez/keyprez/go-src/lib/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCustomerByEmail(email string) (*models.Customer, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, CUSTOMER_COLLECTION)
	var customers []models.Customer

	cursor, err := col.Find(ctx, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &customers); err != nil {
		return nil, err
	}

	if len(customers) > 0 {
		return &customers[0], nil
	}
	return nil, nil
}

func insertCustomer(customer *models.Customer) (*models.Customer, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, CUSTOMER_COLLECTION)

	if _, insertErr := col.InsertOne(ctx, customer); err != nil {
		return nil, insertErr
	}

	return customer, nil
}

func CreateCustomer(email string, StripeID string) (*models.Customer, error) {
	return insertCustomer(&models.Customer{
		ID:       primitive.NewObjectID(),
		StripeID: StripeID,
		Email:    email,
		Created:  time.Now(),
	})
}
