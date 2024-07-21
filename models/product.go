package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name,omitempty" bson:"name,omitempty"`
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
    Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
