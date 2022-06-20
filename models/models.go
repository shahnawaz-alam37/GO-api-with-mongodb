package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Moive string 		  `json:"movie,omitempty"`
	watched bool 		  `json:"watched.omitempty"`
}