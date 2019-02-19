package mongohelpers

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

func RunAggregate(pipeline []bson.M) {

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("feeds").Aggregate(ctx, pipeline)
	catch(err)

	defer curs.Close(ctx)

	elements := []bson.D{}

	for curs.Next(ctx) {
		element := bson.D{}
		err := curs.Decode(&element)
		catch(err)
		elements = append(elements, element)
	}

	return elements
}
