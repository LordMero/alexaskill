package main

import (
	"EllaAlexaSkill/handlers"
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

	r.HandleFunc("/api/feeds/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/feeds/type:{type}&quantity:{quantity}", handlers.InsertDoc).Methods("GET")
	r.HandleFunc("/api/feeds/latest", handlers.GetTotFeed).Methods("GET")

	r.HandleFunc("/api/nappies/", handlers.GetDoc).Methods("GET")
	r.HandleFunc("/api/nappies/type:{type}", handlers.InsertDoc).Methods("GET")

	fmt.Println("Starting server. Listening on port:", port)
	err := http.ListenAndServe("localhost:"+port, r)
	log.Fatal(err)
}
