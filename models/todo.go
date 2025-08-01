package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Item      string             `bson:"item" json:"item"`
	Completed bool               `bson:"completed" json:"completed"`
}
