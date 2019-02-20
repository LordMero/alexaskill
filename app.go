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

	r := mux.NewRouter()
	r.HandleFunc("/api/weights/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/weights/wgt:{wgt}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/api/weights/latest", handlers.GetLatestDoc).Methods("GET")

	r.HandleFunc("/api/feeds/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/feeds/type:{type}&quantity:{quantity}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/api/feeds/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/api/feeds/from:{from}&to:{to}", handlers.GetTotFeed).Methods("GET")

	r.HandleFunc("/api/nappies/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/nappies/type:{type}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/api/nappies/latest", handlers.GetLatestDoc).Methods("GET")
	r.HandleFunc("/api/nappies/from:{from}&to:{to}", handlers.GetTotNappies).Methods("GET")

	r.HandleFunc("/api/baby/", handlers.GetBaby).Methods("GET")

	fmt.Println("Starting server. Listening on port:", port)
	err := http.ListenAndServe("localhost:"+port, r)
	log.Fatal(err)
}
