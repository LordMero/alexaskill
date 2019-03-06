package handlers

import (
	"alexaskill/models"
	"alexaskill/utilities"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func InsertDoc(w http.ResponseWriter, r *http.Request) {

	coll := strings.Split(r.RequestURI, "/")

	fmt.Println(coll[2])

	defer r.Body.Close()

	jb, err := ioutil.ReadAll(r.Body)
	utilities.Catch(err)

	switch coll[2] {
	case "weights":
		wgt := models.Weights{}
		err = json.Unmarshal(jb, &w)
		utilities.Catch(err)

		wgt.Insert()
		_ = json.NewEncoder(w).Encode(wgt)
	case "nappies":
		nps := models.Nappies{}
		err = json.Unmarshal(jb, &nps)
		utilities.Catch(err)

		nps.Insert()

		//fmt.Println(nps)
		_ = json.NewEncoder(w).Encode(nps)
	case "feeds":
		fds := models.Feeds{}
		err = json.Unmarshal(jb, &fds)
		utilities.Catch(err)

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

func GetLatestDoc(w http.ResponseWriter, r *http.Request) {
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

	defer r.Body.Close()


	var dat map[string]string

	jb, err := ioutil.ReadAll(r.Body)
	utilities.Catch(err)
	fmt.Println(string(jb))


	err = json.Unmarshal(jb, &dat)
	utilities.Catch(err)

	fds := models.CountFeeds(dat["when"])

	_ = json.NewEncoder(w).Encode(fds)
	fmt.Fprint(w, fds)
	//fmt.Println(t)
	w.Header().Set("Content-Type", "application/json")
}

func GetTotNappies(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var dat map[string]string

	jb, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(jb, &dat)
	utilities.Catch(err)

	nps := models.CountNappies(dat["when"])

	_ = json.NewEncoder(w).Encode(nps)
	fmt.Fprint(w, nps)
	//fmt.Println(t)
	w.Header().Set("Content-Type", "application/json")
}

func GetBaby(w http.ResponseWriter, r *http.Request) {
	b := models.NewBaby()
	b.Weights = models.GetLastWeight()

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
