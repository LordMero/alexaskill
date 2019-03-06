package main

import (
	"alexaskill/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "5600"

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/weights/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/weights/", handlers.InsertDoc).Methods("POST") // need wgt
	r.HandleFunc("/api/weights/latest", handlers.GetLatestDoc).Methods("GET")

	r.HandleFunc("/api/feeds/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/feeds/", handlers.InsertDoc).Methods("POST") // needs type and quantity
	r.HandleFunc("/api/feeds/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/api/feeds/count", handlers.GetTotFeed).Methods("POST")

	r.HandleFunc("/api/nappies/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/nappies/", handlers.InsertDoc).Methods("POST") // need type
	r.HandleFunc("/api/nappies/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/api/nappies/count", handlers.GetTotNappies).Methods("POST")

	r.HandleFunc("/api/baby/", handlers.GetBaby).Methods("GET")

	fmt.Println("Starting server. Listening on port:", port)
	err := http.ListenAndServe("localhost:"+port, r)
	log.Fatal(err)
}
