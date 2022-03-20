package models

import (
	"time"

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
