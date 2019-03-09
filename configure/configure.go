package configuration

import (
	"context"
	"fmt"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var Db *mongo.Database

const (
	DBNAME     = "babyapp"
	DATELAYOUT = "20060102"
)

func init() {
	gotenv.Load()

	// we use the init function to init mongo db connection

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO")))

	if err != nil {
		log.Panic(err)
	}
	Db = client.Database(DBNAME)
	fmt.Println("connected to mongodb")
}
