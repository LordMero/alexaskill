package handlers

import (
	"alexaskill/models"
	"alexaskill/utilities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InsertDoc(w http.ResponseWriter, r *http.Request) {

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
	w.Header().Set("Content-Type", "application/json")

}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":

		wgt := models.GetAllWeights()
		fmt.Println(wgt)
		_ = json.NewEncoder(w).Encode(wgt)


	case "nappies":
		nps := models.GetAllNappies()
		_ = json.NewEncoder(w).Encode(nps)
		//log.Panic(e)
	case "feeds":
		fds := models.GetAllFeeds()
		_ = json.NewEncoder(w).Encode(fds)
		//log.Panic(e)
	}
}

func GetLatestDoc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":
		wgt := models.GetLastWeight()
		_ = json.NewEncoder(w).Encode(wgt)
		fmt.Println(wgt)
		//log.Panic(e)
	case "nappies":
		nps := models.GetLastNappy()
		fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps)
		//log.Panic(e)
	case "feeds":
		fds := models.GetLastFeed()
		_ = json.NewEncoder(w).Encode(fds)
		fmt.Println(fds)
		//log.Panic(e)
	}

}

func GetTotFeed(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	layout := "2006-01-02T15:04:05Z07:00"

	from, err := time.Parse(layout, params["from"])
	utilities.Catch(err)

	to, err := time.Parse(layout, params["to"])
	utilities.Catch(err)

	fds := models.CountFeeds(from, to)
	_ = json.NewEncoder(w).Encode(fds)
	fmt.Fprint(w, fds)
	//fmt.Println(t)
	w.Header().Set("Content-Type", "application/json")
}

func GetTotNappies(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	layout := "2006-01-02T15:04:05Z07:00"

	from, err := time.Parse(layout, params["from"])
	utilities.Catch(err)

	to, err := time.Parse(layout, params["to"])
	utilities.Catch(err)

	nps := models.CountNappies(from, to)

	_ = json.NewEncoder(w).Encode(nps)
	fmt.Fprint(w, nps)
	//fmt.Println(t)
	w.Header().Set("Content-Type", "application/json")
}

func GetBaby(w http.ResponseWriter, r *http.Request){
	b := models.NewBaby()
	b.Weights =  models.GetLastWeight()

	// handle weights
	bsonBytes, err := bson.Marshal(b.Weights)
	utilities.Catch(err)

	//t := models.Weights{}
	err = bson.Unmarshal(bsonBytes, &b.Weights)
	utilities.Catch(err)


	// handle feeds
	b.Feeds = models.GetLastFeed()
	bsonBytes, err = bson.Marshal(b.Feeds)
	utilities.Catch(err)

	err = bson.Unmarshal(bsonBytes, &b.Feeds)
	utilities.Catch(err)


	// handle nappies
	b.Nappies = models.GetLastNappy()
	bsonBytes, err = bson.Marshal(b.Nappies)
	utilities.Catch(err)

	err = bson.Unmarshal(bsonBytes, &b.Nappies)
	utilities.Catch(err)

	json.NewEncoder(w).Encode(*b)
	w.Header().Set("Content-Type", "application/json")

}

