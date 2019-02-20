package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Baby struct {
	Weights `json:"weights"`
	Feeds `json:"Feeds"`
	Nappies `json:"Nappies"`
}

// define types
type Weights struct {
	Weight    float64   `json:"weight" bson:"weight"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string `json:"-" bson:"-"`
}

type Feeds struct {
	Type      string    `json:"type" bson:"type"`
	Quantity  float64   `json:quantity bson:"quantity"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string`json:"-" bson:"-"`
}

type Nappies struct {
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string `json:"-" bson:"-"`
}

func NewBaby() *Baby {
	return &Baby{
		Weights{Weight: 0, Collection: "weights"},
		Feeds{Type: "", Quantity: 0, Collection: "feeds"},
		Nappies{Type: "", Collection: "nappies"},
	}
}

func main()  {

	b := Baby{
		Weights{Weight: 10, CreatedAt:time.Now(), Collection:"test"},
		Feeds{Type: "bm", CreatedAt:time.Now(), Quantity: 30, Collection:"test2"},
		Nappies{Type: "pee", CreatedAt: time.Now()},
	}
	fmt.Println(b)

	bb := NewBaby()
	fmt.Println(*bb)

	json.NewEncoder(os.Stdout).Encode(bb)
}

