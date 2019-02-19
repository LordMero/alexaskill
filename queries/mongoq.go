package queries

import (
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

func QFoodAt(d time.Time) []bson.M {

	s := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	//s := time.Date(2019, 02, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println("start_date", s)
	e := s.Add(24 * time.Hour)
	fmt.Println("end_date", e)

	q := []bson.M{
		// match
		{"$match": bson.M{"createdat": bson.M{"$gte": s,
			"$lte": e}}},
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
	fmt.Println(q)

	return q

}

/*
func QWeightAt(d time.Time) []bson.M {

}

func QNappiesAt(d time.Time) []bson.M {

}
*/
