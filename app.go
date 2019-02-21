package main

import (
	"alexaskill/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "8080"

func main() {

	b := mux.NewRouter()
	r := b.PathPrefix("/api").Subrouter()
	r.HandleFunc("/weights/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/weights/wgt:{wgt}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/weights/latest", handlers.GetLatestDoc).Methods("GET")

	r.HandleFunc("/feeds/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/feeds/type:{type}&quantity:{quantity}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/feeds/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/feeds/from:{from}&to:{to}", handlers.GetTotFeed).Methods("GET")

	r.HandleFunc("/nappies/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/nappies/type:{type}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/nappies/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/nappies/from:{from}&to:{to}", handlers.GetTotNappies).Methods("GET")

	r.HandleFunc("/baby/", handlers.GetBaby).Methods("GET")

	fmt.Println("Starting server. Listening on port:", port)
	err := http.ListenAndServe("localhost:"+port, r)
	log.Fatal(err)
}
