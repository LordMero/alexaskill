package handlers

import (
	"EllaAlexaSkill/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func InsertDoc(w http.ResponseWriter, r *http.Request) {

	coll := strings.Split(r.RequestURI, "/")

	fmt.Println(coll[2])

	params := mux.Vars(r)
	fmt.Println(params)

	switch coll[2] {
	case "weights":
		p, _ := strconv.ParseFloat(params["wgt"], 64)
		wgts := models.NewWeights(p)
		//wgts.Insert()
		_ = json.NewEncoder(w).Encode(wgts)
	case "nappies":
		nps := models.NewNappies(params["type"])
		//nps.Insert()
		fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps)
	case "feeds":
		q, _ := strconv.ParseFloat(params["quantity"], 64)
		fds := models.NewFeeds(params["type"], q)
		//fds.Insert()
		fmt.Println(fds)
		_ = json.NewEncoder(w).Encode(fds)
	default:
		log.Fatal("fell through")
	}

}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	coll := strings.Split(r.RequestURI, "/")
	switch coll[2] {
	case "weights":
		wgts := models.NewWeights(0)
		_ = json.NewEncoder(w).Encode(wgts.GetAll())
		//fmt.Println(wgts, e)
		//log.Panic(e)
	case "nappies":
		nps := models.NewNappies("")
		_ = json.NewEncoder(w).Encode(nps.GetAll())
		//log.Panic(e)
	case "feeds":
		fds := models.NewFeeds("", 0)
		_ = json.NewEncoder(w).Encode(fds.GetAll())
		//log.Panic(e)
	}
}

func GetTotFeed(w http.ResponseWriter, r *http.Request) {
	fds := models.NewFeeds("", 0)
	t := fds.CountFeeds()
	//_ = json.NewEncoder(w).Encode(t)
	fmt.Fprint(w, t)
}
