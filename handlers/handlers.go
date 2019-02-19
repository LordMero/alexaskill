package handlers

import (
	"EllaAlexaSkill/models"
	"EllaAlexaSkill/utilities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
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

		wgt := models.NewWeights()
		wgt.Weight = p
		wgt.CreatedAt = time.Now()
		wgt.Collection = "weights"

		wgt.Insert()
		_ = json.NewEncoder(w).Encode(wgt)
	case "nappies":
		nps := models.NewNappies()
		nps.Type = params["type"]
		nps.CreatedAt = time.Now()
		nps.Collection = "nappies"

		nps.Insert()

		//fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps)
	case "feeds":
		q, _ := strconv.ParseFloat(params["quantity"], 64)

		fds := models.NewFeeds()

		fds.Type = params["type"]
		fds.Quantity = q
		fds.CreatedAt = time.Now()
		fds.Collection=  "feeds"

		fds.Insert()
		//fmt.Println(fds)
		_ = json.NewEncoder(w).Encode(fds)
	default:
		log.Fatal("fell through")
	}

}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":
		wgt := models.NewWeights()
		_ = json.NewEncoder(w).Encode(wgt.GetAll())
		//fmt.Println(wgts, e)
		//log.Panic(e)
	case "nappies":
		nps := models.NewNappies()
		_ = json.NewEncoder(w).Encode(nps.GetAll())
		//log.Panic(e)
	case "feeds":
		fds := models.NewFeeds()
		_ = json.NewEncoder(w).Encode(fds.GetAll())
		//log.Panic(e)
	}
}

func GetLatestDoc(w http.ResponseWriter, r *http.Request){

	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":
		wgt := models.NewWeights()
		_ = json.NewEncoder(w).Encode(wgt.GetLatest())
		//fmt.Println(wgts, e)
		//log.Panic(e)
	case "nappies":
		nps := models.NewNappies()
		_ = json.NewEncoder(w).Encode(nps.GetLatest())
		//log.Panic(e)
	case "feeds":
		fds := models.NewFeeds()
		_ = json.NewEncoder(w).Encode(fds.GetLatest())
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

	fds := models.NewFeeds()

	t := fds.CountFeeds(from, to)
	//_ = json.NewEncoder(w).Encode(t)
	fmt.Fprint(w, t)
	//fmt.Println(t)
}

func GetTotNappies(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	layout := "2006-01-02T15:04:05Z07:00"

	from, err := time.Parse(layout, params["from"])
	utilities.Catch(err)

	to, err := time.Parse(layout, params["to"])
	utilities.Catch(err)

	fds := models.NewNappies()
	t := fds.CountFeeds(from, to)

	//_ = json.NewEncoder(w).Encode(t)
	fmt.Fprint(w, t)
	//fmt.Println(t)
}

