package models

import (
	"alexaskill/configure"
	"alexaskill/utilities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Nappies struct {
	Type       string `json:"type" bson:"type"`
	CreatedAt  string `json:"created_at" bson:"createdat"`
	Collection string `json:"-" bson:"-"`
}

func NewNappies(t string) *Nappies {
	return &Nappies{
		Type:       t,
		CreatedAt:  time.Now().Format(configuration.DATELAYOUT),
		Collection: "nappies",
	}
}

//Insert document i into coll collection
func (n *Nappies) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("nappies").InsertOne(ctx, n)
	utilities.Catch(err)
}

func GetAllNappies() []Nappies {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("nappies").Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{{"_id", 0}}))
	utilities.Catch(err)

	defer curs.Close(ctx)

	var elements []Nappies

	for curs.Next(ctx) {
		element := Nappies{}
		err := curs.Decode(&element)
		utilities.Catch(err)
		elements = append(elements, element)
	}

	return elements
}

func GetLastNappy() Nappies {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	pipeline := []bson.M{
		{"$limit": 1},
		{"$sort": bson.M{"createdat": -1}},
		{"$project": bson.M{"_id": 0}},
	}

	curs, err := db.Collection("nappies").Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)

	element := Nappies{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		utilities.Catch(err)
	}

	return element

}

func CountNappies(when string) interface{} {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	//pipeline := []bson.M{
	//	// match
	//	{"$match": bson.M{"createdat": bson.M{"$gte": from,
	//		"$lte": to}}},
	//	// group
	//	{"$group": bson.M{
	//		"_id":          bson.M{"type": "$type"}, // "$fieldname" - return the field
	//		"TotalNappies": bson.M{"$sum": 1}}},
	//	// project
	//	{"$project": bson.M{"type": "$_id.type", // project selecte subset of fields
	//		"TotalNappies": "$TotalNappies", // rename fiedls
	//		"_id":          0}},             // 0 means not show _id
	//}

	pipeline := mongo.Pipeline{
		// match
		{{"$match", bson.D{{"createdat", when}}}},
		// group
		{{"$group", bson.D{
			{"_id", bson.D{{"type", "$type"}}}, // "$fieldname" - return the field
			{"totalnappies", bson.D{{"$sum", 1}}},
		}}}, // project
		{{"$project", bson.D{{"type", "$_id.type"}, // project selecte subset of fields
			{"TotalNappies", "$TotalNappies"},      // rename fiedls
			{"_id", 0}}, // 0 means not show _id
		}},
	}

	curs, err := db.Collection("nappies").Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)


	type out struct{
		Type string
		TotalFeeds int32 `bson:"totalfeeds"`
	}

	var e out
	var ee []out

	for curs.Next(ctx) {
		err := curs.Decode(&e)
		utilities.Catch(err)
		ee = append(ee, e)
	}

	return ee
}
