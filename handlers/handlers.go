package handlers

import (
	"alexaskill/models"
	"alexaskill/utilities"
	"encoding/json"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
	"time"
)

func InsertDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers: insert doc invoked")
	coll := strings.Split(r.RequestURI, "/")

	fmt.Println(coll[2])

	params := mux.Vars(r)
	fmt.Println(params)

	switch coll[2] {
	case "weights":
		p, _ := strconv.ParseFloat(params["wgt"], 64)

		wgt := models.NewWeights(p)

		wgt.Insert()
		_ = json.NewEncoder(w).Encode(wgt)
	case "nappies":
		nps := models.NewNappies(params["type"])

		nps.Insert()

		//fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps)
	case "feeds":
		q, _ := strconv.ParseFloat(params["quantity"], 64)

		fds := models.NewFeeds(params["type"], q)

		fds.Insert()
		//fmt.Println(fds)
		_ = json.NewEncoder(w).Encode(fds)
	default:
		log.Fatal("fell through")
	}

}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	coll := strings.Split(r.RequestURI, "/")
	fmt.Println("handlers: get  doc invoked")
	switch coll[2] {
	case "weights":
		wgt := models.NewWeights(0)
		_ = json.NewEncoder(w).Encode(wgt.GetAll())
		//fmt.Println(wgts, e)
		//log.Panic(e)
	case "nappies":
		nps := models.NewNappies("", )
		_ = json.NewEncoder(w).Encode(nps.GetAll())
		//log.Panic(e)
	case "feeds":
		fds := models.NewFeeds("", 0)
		_ = json.NewEncoder(w).Encode(fds.GetAll())
		//log.Panic(e)
	}
}

func GetLatestDoc(w http.ResponseWriter, r *http.Request){
	fmt.Println("handlers: get latest  doc invoked")
	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":
		wgt := models.NewWeights(0)
		_ = json.NewEncoder(w).Encode(wgt.GetLatest())
		fmt.Println(wgt)
		//log.Panic(e)
	case "nappies":
		nps := models.NewNappies("", )
		fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps.GetLatest())
		//log.Panic(e)
	case "feeds":
		fds := models.NewFeeds("", 0)
		_ = json.NewEncoder(w).Encode(fds.GetLatest())
		fmt.Println(fds)
		//log.Panic(e)
	}

}

func GetTotFeed(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handlers: get  tot feeds  invoked")
	params := mux.Vars(r)
	layout := "2006-01-02T15:04:05Z07:00"

	from, err := time.Parse(layout, params["from"])
	utilities.Catch(err)

	to, err := time.Parse(layout, params["to"])
	utilities.Catch(err)

	fds := models.NewFeeds("", 0)

	t := fds.CountFeeds(from, to)
	//_ = json.NewEncoder(w).Encode(t)
	fmt.Fprint(w, t)
	//fmt.Println(t)
}

func GetTotNappies(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handlers: get  tot nappies  invoked")
	params := mux.Vars(r)
	layout := "2006-01-02T15:04:05Z07:00"

	from, err := time.Parse(layout, params["from"])
	utilities.Catch(err)

	to, err := time.Parse(layout, params["to"])
	utilities.Catch(err)

	nps := models.NewNappies("")
	t := nps.CountNappies(from, to)

	//_ = json.NewEncoder(w).Encode(t)
	fmt.Fprint(w, t)
	//fmt.Println(t)
}

func GetBaby(w http.ResponseWriter, r *http.Request){
	b := models.NewBaby()

	// handle weights
	bsonBytes, err := bson.Marshal(b.Weights.GetLatest())
	utilities.Catch(err)

	//t := models.Weights{}
	err = bson.Unmarshal(bsonBytes, &b.Weights)
	utilities.Catch(err)


	// handle feeds
	bsonBytes, err = bson.Marshal(b.Feeds.GetLatest())
	utilities.Catch(err)

	err = bson.Unmarshal(bsonBytes, &b.Feeds)
	utilities.Catch(err)


	// handle nappies
	bsonBytes, err = bson.Marshal(b.Nappies.GetLatest())
	utilities.Catch(err)

	err = bson.Unmarshal(bsonBytes, &b.Nappies)
	utilities.Catch(err)

	json.NewEncoder(w).Encode(*b)
}

