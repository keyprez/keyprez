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

func CreateCustomer(customer *models.Customer, StripeID string) (*models.Customer, error) {
	return insertCustomer(&models.Customer{
		ID:         primitive.NewObjectID(),
		StripeID:   StripeID,
		Email:      customer.Email,
		Name:       customer.Name,
		City:       customer.City,
		Country:    customer.Country,
		Line1:      customer.Line1,
		Line2:      customer.Line2,
		PostalCode: customer.PostalCode,
		State:      customer.State,
		Created:    time.Now(),
	})
}

func UpdateCustomer(id primitive.ObjectID, customer *models.Customer) error {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, CUSTOMER_COLLECTION)

	_, updateErr := col.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"name":       customer.Name,
			"city":       customer.City,
			"country":    customer.Country,
			"line1":      customer.Line1,
			"line2":      customer.Line2,
			"postalCode": customer.PostalCode,
			"state":      customer.State,
		}},
	)

	if updateErr != nil {
		return updateErr
	}

	return nil
}
