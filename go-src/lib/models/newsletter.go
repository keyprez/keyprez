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

func (ns *NewsletterSubscription) IsValid() bool {
	return !ns.ID.IsZero()
}

func GetNewsletterSubscriptionByEmail(email string) (NewsletterSubscription, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return NewsletterSubscription{}, err
	}

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)
	var subscriptions []NewsletterSubscription

	cursor, err := col.Find(ctx, bson.M{"email": email})
	if err != nil {
		return NewsletterSubscription{}, err
	}
	if err = cursor.All(ctx, &subscriptions); err != nil {
		return NewsletterSubscription{}, err
	}

	if len(subscriptions) > 0 {
		return subscriptions[0], nil
	}
	return NewsletterSubscription{}, nil
}

func insertNewsletterSubscription(subscription *NewsletterSubscription) (bool, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)

	if _, insertErr := col.InsertOne(ctx, subscription); err != nil {
		return false, insertErr
	}

	return true, nil
}

func CreateNewsletterSubscription(email string) (bool, error) {
	return insertNewsletterSubscription(&NewsletterSubscription{
		ID:      primitive.NewObjectID(),
		Email:   email,
		Active:  true,
		Created: time.Now(),
	})
}

func UnsubscribeNewsletterSubscription(subscription NewsletterSubscription) (bool, error) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)
	update := bson.M{"$set": bson.M{"active": false}}

	if _, updateErr := col.UpdateByID(ctx, subscription.ID, update); updateErr != nil {
		return false, updateErr
	}

	return true, nil
}
