package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"alexaskill/utilities"
)

type Feeds struct {
	Type       string    `json:"type" bson:"type"`
	Quantity   float64   `json:quantity bson:"quantity"`
	CreatedAt  time.Time `json:"created_at" bson:"createdat"`
	Collection string    `json:"-" bson:"-"`
}

func NewFeeds(t string, q float64) *Feeds {
	return &Feeds{
		Type: t,
		Quantity: q,
		CreatedAt: time.Now(),
		Collection: "feeds",
	}
}

//Insert document i into coll collection
func (n *Feeds) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("feeds").InsertOne(ctx, n)
	utilities.Catch(err)
}

func  GetAllFeeds() []Feeds {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("feeds").Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{{"_id",0}}))
	utilities.Catch(err)

	defer curs.Close(ctx)

	var elements []Feeds

	for curs.Next(ctx) {
		element := Feeds{}
		err := curs.Decode(&element)
		utilities.Catch(err)
		elements = append(elements, element)
	}

	return elements
}

func  GetLastFeed() Feeds {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	pipeline := []bson.M{
		{"$limit": 1},
		{"$sort": bson.M{"createdat": -1}},
		{"$project": bson.M{"_id": 0}},
	}

	curs, err := db.Collection("feeds").Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)

	element := Feeds{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		utilities.Catch(err)
	}

	return element



}

func CountFeeds(from time.Time, to time.Time) Feeds {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	pipeline := []bson.M{
		// match
		{"$match": bson.M{"createdat": bson.M{"$gte": from,
			"$lte": to}}},
		// group
		{"$group": bson.M{
			"_id":        bson.M{"type": "$type"}, // "$fieldname" - return the field
			"TotalFeeds":  bson.M{"$sum": 1}}},
		// project
		{"$project": bson.M{"type": "$_id.type", // project selecte subset of fields
			"TotalFeeds":  "$TotalFeeds", // rename fiedls
			"_id":        0}}, // 0 means not show _id
	}

	curs, err := db.Collection("feeds").Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)

	element := Feeds{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		utilities.Catch(err)
	}

	return element
}


