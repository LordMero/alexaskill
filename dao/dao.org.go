package configuration

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var Db *mongo.Database

const (
	CONSTRING =  "mongodb://localhost:27017"
	DBNAME = "babyapp"
	)

func init()  {
	// we use the init function to init mongo db connection
	c, e := mongo.Connect(context.Background(), CONSTRING)
	if e != nil {log.Fatal(e)}

	Db = c.Database(DBNAME)

}
