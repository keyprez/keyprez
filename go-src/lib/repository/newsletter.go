package repository

import (
	"context"
	"time"

	"github.com/keyprez/keyprez/go-src/lib/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNewsletterSubscriptionByEmail(email string) (*models.NewsletterSubscription, error) {
	ctx := context.TODO()
	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	col := GetMongoCollection(mongoClient, NEWSLETTER_COLLECTION)
	var subscriptions []models.NewsletterSubscription

	cursor, err := col.Find(ctx, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &subscriptions); err != nil {
		return nil, err
	}

	if len(subscriptions) > 0 {
		return &subscriptions[0], nil
	}
	return nil, nil
}

func insertNewsletterSubscription(subscription *models.NewsletterSubscription) (bool, error) {
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
	return insertNewsletterSubscription(&models.NewsletterSubscription{
		ID:      primitive.NewObjectID(),
		Email:   email,
		Active:  true,
		Created: time.Now(),
	})
}

func UnsubscribeNewsletterSubscription(subscription *models.NewsletterSubscription) (bool, error) {
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
