// models
package models

import (
	"alexaskill/controllers"
	"alexaskill/utilities"
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

func (w Weights) GetAll() Weights {

	sb, err := bson.Marshal(controllers.GetAll(w.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, w)

	return w

}

func (w Weights) GetLatest() Weights{

	sb, err := bson.Marshal(controllers.GetLatest(w.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, w)

	return w
}


//  FEEDS

func (f Feeds) Insert() {
	controllers.InsertOne(f, f.Collection)
}

func (f Feeds) GetAll() Feeds {
	sb, err := bson.Marshal(controllers.GetAll(f.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, f)

	return f
}

func (f Feeds) GetLatest() Feeds {
	sb, err := bson.Marshal(controllers.GetLatest(f.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, f)

	return f
}

func (f Feeds) CountFeeds(from time.Time, to time.Time) Feeds {
	sb, err := bson.Marshal(controllers.CountFeeds(from, to))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, f)

	return f
}

// NAPPIES


func (n Nappies) Insert(){
	controllers.InsertOne(n, n.Collection)
}

func (n Nappies) GetAll() Nappies {
	sb, err := bson.Marshal(controllers.GetAll(n.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, n)

	return n
}

func (n Nappies) GetLatest() Nappies {

	sb, err := bson.Marshal(controllers.GetLatest(n.Collection))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, n)

	return n
}

func (n Nappies) CountNappies(from time.Time, to time.Time) Nappies {

	sb, err := bson.Marshal(controllers.CountNappies(from, to))
	utilities.Catch(err)

	_ = bson.Unmarshal(sb, n)

	return n
}



