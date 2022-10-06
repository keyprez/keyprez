package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StripeID   string             `bson:"stripeid,omitempty"`
	Email      string             `bson:"email,omitempty"`
	Name       string             `bson:"name,omitempty"`
	City       string             `bson:"city,omitempty"`
	Country    string             `bson:"country,omitempty"`
	Line1      string             `bson:"line1,omitempty"`
	Line2      string             `bson:"line2,omitempty"`
	PostalCode string             `bson:"postalCode,omitempty"`
	State      string             `bson:"state,omitempty"`
	Created    time.Time          `bson:"created,omitempty"`
}

func (ns *Customer) IsValid() bool {
	return !ns.ID.IsZero()
}
