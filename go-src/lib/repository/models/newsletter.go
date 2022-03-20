package models

import (
	"time"

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
