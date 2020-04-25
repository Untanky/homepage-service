package models

import (
	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type content struct {
	Type string `json:"type" bson:"type"`

	Text string `json:"text,omitempty" bson:"text,omitempty"`

	Title string `json:"title,omitempty" bson:"title,omitempty"`

	VideoId string `json:"videoId,omitempty" bson:"videoId,omitempty"`

	Source        string `json:"src,omitempty" bson:"src,omitempty"`
	AlternateText string `json:"alt,omitempty" bson:"alt,omitempty"`

	Link               string `json:"link,omitempty" bson:"link,omitempty"`
	ImageUrl           string `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	ImageAlternateText string `json:"imageAltText,omitempty" bson:"imageAltText,omitempty"`

	Language string `json:"language,omitempty" bson:"language,omitempty"`
	Code string `json:"code,omitempty" bson:"code,omitempty"`

	Caption string `json:"caption,omitempty" bson:"caption,omitempty"`
}

type section struct {
	Layout  string    `json:"layout" bson:"layout"`
	Content []content `json:"content" bson:"content"`
}

type Post struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel                        `bson:",inline"`
	Id                int                   `json:"id" bson:"id"`
	Title             string                `json:"title" bson:"title"`
	Author            string                `json:"author" bson:"author"`
	Created	          primitive.DateTime    `json:"created" bson:"created"`
	Updated           primitive.DateTime    `json:"updated" bson:"updated"`
	Section           []section             `json:"section" bson:"section"`
}
