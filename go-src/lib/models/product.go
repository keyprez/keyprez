package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This will be a product struct when we add such a collection.
// Test data only for now
type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Date    time.Time          `bson:"date,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Text    string             `bson:"text,omitempty"`
	MovieId primitive.ObjectID `bson:"movie_id,omitempty"`
}

func GetProducts() ([]Comment, error) {
	ctx := context.TODO()
	productCollection := GetMongoCollection("comments")
	var comments []Comment

	selectOpts := options.Find()
	selectOpts.SetLimit(3)

	cursor, err := productCollection.Find(ctx, bson.M{"email": "john_bishop@fakegmail.com"}, selectOpts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
