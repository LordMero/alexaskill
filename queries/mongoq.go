package queries

import (
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

func TodayFood(d time.Time) []bson.M {
	q := []bson.M{bson.M{
		"$match": bson.M{
			"createdat": bson.M{
				"$gte": "2019-02-17T00:00:00.000+0000",
				"$lte": "2019-02-18T00:00:00.000+0000",
			},
		},
	},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"type":     "$type",
					"quantity": "$quantity",
				},
				"COUNT(id)": bson.M{
					"$sum": 1,
				},
				"SUM(quantity)": bson.M{
					"$sum": "$quantity",
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"COUNT(id)":     "$COUNT(id)",
				"type":          "$_id.type",
				"SUM(quantity)": "$SUM(quantity)",
				"_id":           0,
			},
		},
	}
	return q
}

func TodayFoodv2(d time.Time) []bson.M {

	s := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println("start_date", s)
	e := s.Add(24 * time.Hour)
	fmt.Println("end_date", e)

	q := []bson.M{
		{"$match": bson.M{"createdat": bson.M{"$gte": s,
			"$lte": e}}},
		{"$group": bson.M{
			"_id":        bson.M{"type": "$type", "quantity": "$quantity"},
			"TotalFeed":  bson.M{"$sum": 1},
			"TotalQuant": bson.M{"$sum": "$quantity"}}},
		{"$project": bson.M{"TotalFeed": "$TotalFeed",
			"type":       "$_id.type",
			"TotalQuant": "$TotalQuant"}},
	}

	return q

}
