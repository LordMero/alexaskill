// models
package models

import (
	"EllaAlexaSkill/controllers"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type BabyCollections mongo.Collection

type Baby struct {
	Weights `json:"weights"`
	Feeds `json:"Feeds"`
	Nappies `json:"Nappies"`
}

// define types
type Weights struct {
	Weight    float64   `json:"weight" bson:"weight"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string `json:"-" bson:"-"`
}

type Feeds struct {
	Type      string    `json:"type" bson:"type"`
	Quantity  float64   `json:quantity bson:"quantity"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string`json:"-" bson:"-"`
}

type Nappies struct {
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
	Collection string `json:"-" bson:"-"`
}

type totalfeed struct {
	c int     `json:"count" bson:"TotalFeed"`
	t string  `json:"type" bson:"type"`
	q float64 `json:"totquantity" bson:"TotalQuant"`
}

// ================= implement constructors ===========================
func NewWeights(w float64) *Weights {
	return &Weights{
		Weight: w,
		CreatedAt: time.Now(),
		Collection: "weights",
	}

}

func NewFeeds(t string, q float64) *Feeds {
	return &Feeds{
		Type: t,
		Quantity: q,
		CreatedAt: time.Now(),
		Collection: "feeds",
	}
}

func NewNappies(t string) *Nappies {
	return &Nappies{
		Type: t,
		CreatedAt: time.Now(),
		Collection: "nappies",
	}
}

func NewBaby() *Baby {
	return &Baby{
		Weights{Weight:0, Collection: "weights"},
		Feeds{Type:"", Quantity:0, Collection: "feeds"},
		Nappies{Type:"", Collection: "nappies"},
	}

}

//Insert document i into coll collection
func (w Weights) Insert() {
	controllers.InsertOne(w, w.Collection)
}

func (w Weights) GetAll() bson.D {
	return controllers.GetAll(w.Collection)
}

func (w Weights) GetLatest() bson.D{
	return controllers.GetLatest(w.Collection)
}


//  FEEDS

func (f Feeds) Insert() {
	controllers.InsertOne(f, f.Collection)
}

func (f Feeds) GetAll() bson.D {
	return controllers.GetAll(f.Collection)
}

func (f Feeds) GetLatest() bson.D {
	return controllers.GetLatest(f.Collection)
}

func (f Feeds) CountFeeds(from time.Time, to time.Time) bson.D {
	return controllers.CountFeeds(from, to)
}

// NAPPIES


func (n Nappies) Insert(){
	controllers.InsertOne(n, n.Collection)
}

func (n Nappies) GetAll() bson.D {
	return controllers.GetAll(n.Collection)
}

func (n Nappies) GetLatest() bson.D {
	return controllers.GetLatest(n.Collection)
}

func (n Nappies) CountNappies(from time.Time, to time.Time) bson.D {
	return controllers.CountNappies(from, to)
}



