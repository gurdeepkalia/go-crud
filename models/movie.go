package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"movie_name,omitempty" bson:"movie_name,omitempty" validate:"required"`
	Watched bool               `json:"watched" bson:"watched"`
}
