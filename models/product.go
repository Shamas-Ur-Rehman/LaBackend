package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Price       float64            `json:"price" binding:"required"`
	Inventory   int                `json:"inventory" binding:"required"`
	ImageURL    string             `json:"imageUrl,omitempty"`

	Category string   `json:"category" binding:"required"`
	THC      string   `json:"thc"`
	CBD      string   `json:"cbd"`
	Strain   string   `json:"strain"`
	Effects  []string `json:"effects"`
	Badge    string   `json:"badge,omitempty"`
}
