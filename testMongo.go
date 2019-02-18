package main

import (
	"EllaAlexaSkill/models"
	"context"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"os"
)


const CONNECTIONSTRING = "mongodb://localhost:27017"
const DBNAME = "babyapp"

// context are used to manage connection to servers,
// within the context we can specify for instance how long the connetion can be open
// background is an empty context used for testing

var db *mongo.Database

func init()  {

	// create a client
	client, err := mongo.NewClient(CONNECTIONSTRING)
	must(err)

	// connect to the client
	err = client.Connect(context.Background())
	must(err)

	// get database from the client
	db = client.Database(DBNAME)

}


func main()  {
	//var b models.Baby
	//var xb []models.Baby

	u := db.Collection("baby")

	/* create a fake baby
	b := models.Baby{
		Weight: 4.16,
		NappyWee: 1,
		NappyPoo: 0,
		FeedQ: 1,
		FeedN: 140,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	*/
	//_, err = u.InsertOne(context.Background(), b)
	//must(err)

	cur, err := u.Find(context.Background(), bson.D{})
	must(err)

	defer cur.Close(context.Background())

	var res models.Baby
	var xb []models.Baby

	for cur.Next(context.Background()) {

		err := cur.Decode(&res)
		if err != nil { log.Fatal(err) }
		xb = append(xb, res)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(os.Stdout).Encode(xb)
	//fmt.Println(res)

}

func must(e error)  {
	if e != nil {
		log.Fatal(e)
	}
}

