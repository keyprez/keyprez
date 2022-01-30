package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	StripeID string             `bson:"stripeid,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Created  time.Time          `bson:"created,omitempty"`
}

func (ns *Customer) IsValid() bool {
	return !ns.ID.IsZero()
}

func GetCustomerByEmail(email string) (Customer, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return Customer{}, err
	}

	col := GetMongoCollection(mongoClient, CUSTOMER_COLLECTION)
	var customers []Customer

	cursor, err := col.Find(ctx, bson.M{"email": email})
	if err != nil {
		return Customer{}, err
	}
	if err = cursor.All(ctx, &customers); err != nil {
		return Customer{}, err
	}

	if len(customers) > 0 {
		return customers[0], nil
	}
	return Customer{}, nil
}

func insertCustomer(customer *Customer) (Customer, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return Customer{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, CUSTOMER_COLLECTION)

	if _, insertErr := col.InsertOne(ctx, customer); err != nil {
		return Customer{}, insertErr
	}

	return *customer, nil
}

func CreateCustomer(email string, StripeID string) (Customer, error) {
	return insertCustomer(&Customer{
		ID:       primitive.NewObjectID(),
		StripeID: StripeID,
		Email:    email,
		Created:  time.Now(),
	})
}
