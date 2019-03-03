package models

import (
	"alexaskill/configure"
	"alexaskill/utilities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// define types
type Weights struct {
	Weight     float64 `json:"weight" bson:"weight"`
	CreatedAt  string  `json:"created_at" bson:"createdat"`
	Collection string  `json:"-" bson:"-"`
}

var db = configuration.Db

func NewWeights(w float64) *Weights {
	return &Weights{
		Weight:     w,
		CreatedAt:  time.Now().Format(configuration.DATELAYOUT),
		Collection: "weights",
	}

}

//Insert document i into coll collection
func (w *Weights) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("weights").InsertOne(ctx, w)
	utilities.Catch(err)
}

func GetAllWeights() []Weights {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("weights").Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{{"_id", 0}}))
	utilities.Catch(err)

	defer curs.Close(ctx)

	var elements []Weights

	for curs.Next(ctx) {
		element := Weights{}
		err := curs.Decode(&element)
		utilities.Catch(err)
		elements = append(elements, element)
	}

	return elements
}

func GetLastWeight() Weights {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	pipeline := []bson.M{
		{"$limit": 1},
		{"$sort": bson.M{"createdat": -1}},
		{"$project": bson.M{"_id": 0}},
	}

	curs, err := db.Collection("weights").Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)

	element := Weights{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		utilities.Catch(err)
	}

	return element

}
