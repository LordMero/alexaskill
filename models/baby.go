package models

import (
	"time"
)

type Baby struct {
	Weights
	Feeds
	Nappies
}

// define types
type Weights struct {
	Weight    float64   `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

type Feeds struct {
	Type      string    `json:"type"`
	Quantity  float64   `json:quantity`
	CreatedAt time.Time `json:"created_at"`
}

type Nappies struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type totalfeed struct {
	c int     `json:"count" bson:"TotalFeed"`
	t string  `json:"type" bson:"type"`
	q float64 `json:"totquantity" bson:"TotalQuant"`
}
