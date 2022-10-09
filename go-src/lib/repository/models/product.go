package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       int32              `bson:"price,omitempty" json:"price,omitempty"`
	PriceID     string             `bson:"price_id,omitempty" json:"priceId,omitempty"`
	Stock       int32              `bson:"stock,omitempty" json:"stock,omitempty"`
	Slug        string             `bson:"slug,omitempty" json:"slug,omitempty"`
	Active      bool               `bson:"active,omitempty" json:"active,omitempty"`
}
