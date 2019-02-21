// all mongodb operations are written here
package controllers

import (
	"alexaskill/dao"
	"alexaskill/utilities"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	"fmt"
)


// import db from configuration
var db = configuration.Db

//Insert document i into coll collection
func InsertOne(i interface{}, coll string) {
	fmt.Println("controllers: insert one invoked")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection(coll).InsertOne(ctx, i)
	utilities.Catch(err)

}

//GetAll get all the documents from collection coll
func  GetAll(coll string) bson.D {
	fmt.Println("controllers: get all invoked")
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection(coll).Find(ctx, bson.D{})
	utilities.Catch(err)

	defer curs.Close(ctx)

	elements := bson.D{}

	for curs.Next(ctx) {
		err := curs.Decode(&elements)
		utilities.Catch(err)
	}

	return elements
}

// Get the latest entry in the the collection coll
func GetLatest(coll string) bson.D {
	fmt.Println("controllers: get latest invoked")
	pipeline := []bson.M{
		{"$limit": 1},
		{"$sort": bson.M{"createdat": -1}},
		{"$project": bson.M{"_id": 0}},
	}

	return runAggregate(pipeline, db.Collection(coll))
}

// Count how many feeds for any give 2 points in time
func CountFeeds(from time.Time, to time.Time) bson.D {
	fmt.Println("controllers: CountFeeds invoked")
	pipeline := []bson.M{
		// match
		{"$match": bson.M{"createdat": bson.M{"$gte": from,
			"$lte": to}}},
		// group
		{"$group": bson.M{
			"_id":        bson.M{"type": "$type"}, // "$fieldname" - return the field
			"TotalFeed":  bson.M{"$sum": 1},
			"TotalQuant": bson.M{"$sum": "$quantity"}}},
		// project
		{"$project": bson.M{"type": "$_id.type", // project selecte subset of fields
			"TotalFeed":  "$TotalFeed", // rename fiedls
			"TotalQuant": "$TotalQuant",
			"_id":        0}}, // 0 means not show _id
	}

	return runAggregate(pipeline, db.Collection("feeds"))
}

// Count how many nappies for any give 2 points in time
func CountNappies(from time.Time, to time.Time) bson.D {

	pipeline := []bson.M{
		// match
		{"$match": bson.M{"createdat": bson.M{"$gte": from,
			"$lte": to}}},
		// group
		{"$group": bson.M{
			"_id":        bson.M{"type": "$type"}, // "$fieldname" - return the field
			"TotalNappies":  bson.M{"$sum": 1}}},
		// project
		{"$project": bson.M{"type": "$_id.type", // project selecte subset of fields
			"TotalNappies":  "$TotalNappies", // rename fiedls
			"_id":        0}}, // 0 means not show _id
	}

	return runAggregate(pipeline, db.Collection("nappies"))
}


func runAggregate(pipeline []bson.M, coll *mongo.Collection) bson.D{

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := coll.Aggregate(ctx, pipeline)
	utilities.Catch(err)

	defer curs.Close(ctx)

	elements := bson.D{}

	for curs.Next(ctx) {
		//element := bson.D{}
		err := curs.Decode(&elements)
		utilities.Catch(err)
	}

	return elements
}
