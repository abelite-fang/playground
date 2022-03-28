package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	//BoardId  primitive.ObjectID `json:"board,omitempty" validate:required`
	Title    string `json:"title,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Content  string `json:"content,omitempty" validate:"required"`
}

type Board struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
}
