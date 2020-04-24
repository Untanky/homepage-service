package models

import (
	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel                     `bson:",inline"`
	Title             string             `json:"title" bson:"title"`
	Author            string             `json:"author" bson:"author"`
	Created	          primitive.DateTime `json:"created" bson:"created"`
	Updated           primitive.DateTime `json:"updated" bson:"updated"`
}
