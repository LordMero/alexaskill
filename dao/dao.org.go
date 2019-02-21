package configuration

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"fmt"
)

var Db *mongo.Database

const (
	CONSTRING =  "mongodb://localhost:27017"
	DBNAME = "babyapp"
	)

func init()  {
	// we use the init function to init mongo db connection
	//c, e := mongo.Connect(context.Background(), CONSTRING)
	c, e  := mongo.NewClient(options.Client().ApplyURI(CONSTRING))
	if e != nil {log.Fatal(e)}
	Db = c.Database(DBNAME)
	fmt.Println("connected to mongodb")
}
