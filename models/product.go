package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Stock int                `json:"stock" bson:"stock"`
	Price float64            `json:"price" bson:"price"`
}
