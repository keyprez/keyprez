package models

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StripeID   string             `bson:"stripeid,omitempty"`
	Email      string             `bson:"email,omitempty" validate:"required,email"`
	Name       string             `bson:"name,omitempty" validate:"required,min=2"`
	City       string             `bson:"city,omitempty" validate:"required"`
	Country    string             `bson:"country,omitempty" validate:"required,eq=NO|eq=SE|eq=DK"`
	Line1      string             `bson:"line1,omitempty" validate:"required"`
	Line2      string             `bson:"line2,omitempty" validate:"omitempty,min=2"`
	PostalCode string             `bson:"postalCode,omitempty" validate:"required"`
	State      string             `bson:"state,omitempty"`
	Created    time.Time          `bson:"created,omitempty"`
}

func (ns *Customer) HasValidID() bool {
	return !ns.ID.IsZero()
}

func (customer *Customer) IsValid() bool {
	validate := validator.New()
	err := validate.Struct(customer)
	return err == nil
}
