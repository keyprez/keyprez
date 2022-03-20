package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int32              `bson:"price,omitempty"`
	PriceID     string             `bson:"price_id,omitempty"`
	Stock       int32              `bson:"stock,omitempty"`
	Slug        string             `bson:"slug,omitempty"`
	Active      bool               `bson:"active,omitempty"`
}
