
package controllers

type BabyController {
	Insert
	GetAll 
	GetLatest
	CountFeeds
	CountNappies
}


// import db from configuration
var db = configuration.Db

func Insert(w interface{}, c string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection(c).InsertOne(ctx, w)
	catch(err)

}

func GetAll(c string) []bson.M {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection(c).Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	elements := []bson.M{}

	for curs.Next(ctx) {
		element := Weights{}
		err := curs.Decode(&element)
		catch(err)
		elements = append(elements, element)
	}

	return elements
}

func GetLatest(c string) []bson.D {
	pipeline := queries.QLatest()

	return mongohelpers.RunAggregate(pipeline, c)
}

func CountFeeds() []bson.D {
	pipeline := queries.QFoodAt(time.Now())

	return mongohelpers.RunAggregate(pipeline)
}

func CountNappies() []bson.D {

	pipeline := queries.QNappiesAt(time.Now())

	return mongohelpers.RunAggregate(pipeline)
}
