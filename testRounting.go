package main

import (
	"EllaAlexaSkill/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

var b models.Baby

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHand)
	r.HandleFunc("/add", AddHand)

	http.ListenAndServe(":8080", r)
}

func HomeHand(w http.ResponseWriter, r *http.Request)  {
	// encode the structure into a json and write to the http writer
	e := json.NewEncoder(w).Encode(b)
	if e != nil {
		log.Fatal(e)
	}
	json.NewEncoder(os.Stdout).Encode(b)
}

func AddHand(w http.ResponseWriter, r *http.Request){
	b.NappyPoo = 0
	b.NappyWee = 1
	b.CreatedAt = time.Now()

}