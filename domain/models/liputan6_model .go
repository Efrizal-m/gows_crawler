package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Liputan6News struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Headline string             `json:"headline,omitempty" validate:"required"`
	URL      string             `json:"url,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" validate:"required"`
	Summary  string             `json:"summary,omitempty" validate:"required"`
}

type Liputan6NewsData struct {
	Data []Liputan6News
}
