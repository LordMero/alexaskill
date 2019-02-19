package tests

import (
	"fmt"
	"net/http"
	"time"
)

type baby struct {
	Id time.Time `json:id`
	Weight float64 `json:"weight"`
	NappyWee int `json:"nappy_wee"`
	NappyPoo int `json:"nappy_poo"`
	FeedQ float64 `json:"feed_q"`
	FeedN float64 `json:"feed_n"`
}

func (b baby) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println("We are inside ServeHttp")
}

func main() {
	var h baby
	http.ListenAndServe(":8080", h)
}
