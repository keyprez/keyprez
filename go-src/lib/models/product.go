package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Date    time.Time          `bson:"date,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Text    string             `bson:"text,omitempty"`
	MovieId primitive.ObjectID `bson:"movie_id,omitempty"`
}
