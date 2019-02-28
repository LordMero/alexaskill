package configuration

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"fmt"
	"context"
	"time"
)

var Db *mongo.Database

const (
	CONSTRING =  "mongodb://localhost:27017"
	DBNAME = "babyapp"
)

func init()  {
	// we use the init function to init mongo db connection
	//c, e := mongo.Connect(context.Background(), CONSTRING)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONSTRING))

	if err != nil {log.Panic(err)}
	Db = client.Database(DBNAME)
	fmt.Println("connected to mongodb")
}