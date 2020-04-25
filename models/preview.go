package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Preview struct {
	Id int `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Created primitive.DateTime `json:"created" bson:"created"`
	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Text string `json:"text" bson:"text"`
	PostUrl string `json:"postUrl" bson:"postUrl"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	ImageAltText string `json:"imageAlt" bson:"imageAlt"`
	DataUrl string `json:"dataUrl" bson:"dataUrl"`
}