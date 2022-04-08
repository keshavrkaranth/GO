package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Movie   string             `json:"movie,omitempty" `
	Watched bool               `json:"watched,omitempty" `
}
