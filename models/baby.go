// models
package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type BabyCollections mongo.Collection

type Baby struct {
	Weights `json:"weights"`
	Feeds `json:"Feeds"`
	Nappies `json:"Nappies"`
}

type totalfeed struct {
	c int     `json:"count" bson:"TotalFeed"`
	t string  `json:"type" bson:"type"`
	q float64 `json:"totquantity" bson:"TotalQuant"`
}

func NewBaby() *Baby {
	return &Baby{
		Weights{Weight:0, Collection: "weights"},
		Feeds{Type:"", Quantity:0, Collection: "feeds"},
		Nappies{Type:"", Collection: "nappies"},
	}

}


