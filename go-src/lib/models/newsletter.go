package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsletterSubscription struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Active  bool               `bson:"active,omitempty"`
	Created time.Time          `bson:"created,omitempty"`
}

func GetNewsletterSubscriptionByEmail(email string) ([]NewsletterSubscription, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)
	var subscriptions []NewsletterSubscription

	cursor, err := col.Find(ctx, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func insertNewsletterSubscription(subscription NewsletterSubscription) {

}

func CreateNewsletterSubscription(email string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	mongoClient, err := GetMongoClient()
	if err != nil {
		return false, err
	}

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)
	newNewsletterSubscription := &NewsletterSubscription{
		ID:      primitive.NewObjectID(),
		Email:   email,
		Active:  true,
		Created: time.Now(),
	}

	_, insertErr := col.InsertOne(ctx, newNewsletterSubscription)

	if insertErr != nil {
		return false, insertErr
	}

	return true, nil
}
