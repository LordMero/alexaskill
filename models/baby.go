package models

import (
	configuration "EllaAlexaSkill/dao"
	"EllaAlexaSkill/mongohelpers"
	"EllaAlexaSkill/queries"
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// import db from configuration
var db = configuration.Db

type Baby struct {
	Weights
	Feeds
	Nappies
}

// define types
type Weights struct {
	Weight    float64   `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

type Feeds struct {
	Type      string    `json:"type"`
	Quantity  float64   `json:quantity`
	CreatedAt time.Time `json:"created_at"`
}

type Nappies struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type totalfeed struct {
	c int     `json:"count" bson:"TotalFeed"`
	t string  `json:"type" bson:"type"`
	q float64 `json:"totquantity" bson:"TotalQuant"`
}

// ================ Weights Methods
func NewWeights(wgt float64) *Weights {
	return &Weights{
		Weight:    wgt,
		CreatedAt: time.Now(),
	}
}

func (w Weights) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("weights").InsertOne(ctx, w)
	catch(err)

}

func (w Weights) GetAll() []Weights {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("weights").Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	elements := []Weights{}

	for curs.Next(ctx) {
		element := Weights{}
		err := curs.Decode(&element)
		catch(err)
		elements = append(elements, element)
	}

	return elements
}

//func (w Weights) GetLatest() Weights {
//
//}

// ================ Feeds Methods
func NewFeeds(t string, q float64) *Feeds {
	return &Feeds{
		Type:      t,
		Quantity:  q,
		CreatedAt: time.Now(),
	}
}

// implement operation on db
func (f Feeds) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("feeds").InsertOne(ctx, f)
	catch(err)

}

func (f Feeds) GetAll() []Feeds {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("feeds").Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	elements := []Feeds{}

	for curs.Next(ctx) {
		element := Feeds{}
		err := curs.Decode(&element)
		catch(err)
		elements = append(elements, element)
	}

	return elements
}

func (f Feeds) GetLatest() Feeds {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("feeds").Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	element := Feeds{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		catch(err)
	}

	return element
}

func (f Feeds) CountFeeds() []bson.D {

	pipeline := queries.QFoodAt(time.Now())

	return mongohelpers.RunAggregate(pipeline)
}

func NewNappies(t string) *Nappies {
	return &Nappies{
		Type:      t,
		CreatedAt: time.Now(),
	}
}

// implement operation on db
func (n Nappies) Insert() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := db.Collection("nappies").InsertOne(ctx, n)
	catch(err)

}

func (n Nappies) GetAll() []Nappies {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("nappies").Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	elements := []Nappies{}

	for curs.Next(ctx) {
		element := Nappies{}
		err := curs.Decode(&element)
		catch(err)
		elements = append(elements, element)
	}

	return elements
}

func (n Nappies) GetLatest() Nappies {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	curs, err := db.Collection("nappies").Find(ctx, bson.D{})
	catch(err)

	defer curs.Close(ctx)

	element := Nappies{}

	for curs.Next(ctx) {
		err := curs.Decode(&element)
		catch(err)
	}

	return element
}

func catch(e error) {
	if e != nil {
		log.Panic(e)
	}
}
