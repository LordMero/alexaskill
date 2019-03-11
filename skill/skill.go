package main

import (
	mod "alexaskill/models"
	amod "alexaskill/skill/models"
	"alexaskill/utilities"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var w []mod.Weights
var f []mod.Feeds
var n []mod.Nappies

/*func apihandler(i interface{}, resp *http.Response){

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	switch i.(type){
	case []mod.Weights:
		fmt.Println("Weights")
		_ = json.Unmarshal(b, &w)
	case []mod.Feeds:
		fmt.Println("Feeds")
		_ = json.Unmarshal(b, &f)
	case []mod.Nappies:
		fmt.Println("Nappies")
		_ = json.Unmarshal(b, &n)
	}


}*/

func apihandler(resp *http.Response) map[string]interface{} {

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))

	var o map[string]interface{}
	err := json.Unmarshal(b, &o)
	utilities.Catch(err)

	return o

}

// TODO: - COMPLETE NAPPY COUNT AND FEED COUNT

func WriteRequest(i interface{}, coll string) {

	url := os.Getenv("WEBSITE") + coll + "/"

	sb, err := json.Marshal(i)
	utilities.Catch(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sb))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func HandleRequest(ctx context.Context, i amod.AlexaComplexRequest) (amod.AlexaResponse, error) {
	// Use Spew to output the request for debugging purposes:
	fmt.Println("---- Dumping Input Map: ----")
	spew.Dump(i)
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	log.Printf("Request type is ", i.Request.Intent.Name)

	// Create a response object
	resp := amod.CreateResponse()

	// Customize the response for each Alexa Intent
	switch i.Request.Intent.Name {
	case "AddNappy":
		t := i.Request.Intent.Slots.Kind.Value
		s := fmt.Sprintf("I am adding a %s nappy. I hope it doesn't smell!", t)

		nn := mod.NewNappies(t)

		WriteRequest(nn, "nappies")

		resp.Say(s)
	case "AddFeed":
		t := i.Request.Intent.Slots.Type.Value
		q, _ := strconv.ParseFloat(i.Request.Intent.Slots.Quantity.Value, 64)
		s := fmt.Sprintf("I am adding %s of %s. Thanks", q, t)

		nf := mod.NewFeeds(t, q)

		WriteRequest(nf, "feeds")

		resp.Say(s)
	case "AddBreastFeed":
		/*logType := i.Request.Intent.Slots.LogType.Value
		dur := i.Request.Intent.Slots.Duration.Value
		timeUnit := i.Request.Intent.Slots.TimeUnit.Value
		s := fmt.Sprintf("I am %sing a breast feeding session of %s %s.", logType, dur, timeUnit)
		if timeUnit == "hour" {

			df, _ := strconv.ParseFloat(dur, 64)
			dur = fmt.Sprintf("%f", df*60)

			}
		u := fmt.Sprintf(webstr + "feeds/type:breast&quantity:%s", dur)
		hr, _ := http.Get(u)
		fmt.Println(hr)
		resp.Say(s)*/
	case "AddWeight":
		n := i.Request.Intent.Slots.Name.Value
		q, _ := strconv.ParseFloat(i.Request.Intent.Slots.Wgt.Value, 64)
		s := fmt.Sprintf("%s now weights %s. Got it!", n, q)

		nw := mod.NewWeights(q)

		WriteRequest(nw, "weights")

		resp.Say(s)
	case "GetWeight":
		url := os.Getenv("WEBSITE") + "weights/latest"
		hr, _ := http.Get(url)

		o := apihandler(hr)

		n := i.Request.Intent.Slots.Name.Value
		q := o["weight"].(float64)
		s := fmt.Sprintf("%s weights %f kilos!", n, q)
		resp.Say(s)
	case "GetLastFeedTime":
		n := i.Request.Intent.Slots.Name.Value
		q := "12"
		s := fmt.Sprintf("%s lastly ate at %s", n, q)
		resp.Say(s)
	case "GetLastFeedQuantity":
		url := os.Getenv("WEBSITE") + "feeds/latest"
		hr, _ := http.Get(url)

		o := apihandler(hr)
		n := i.Request.Intent.Slots.Name.Value
		q := o["Quantity"].(float64)
		t := o["type"].(string)
		s := fmt.Sprintf("Last feed %s had %f of %s", n, q, t)
		resp.Say(s)
	case "GetLastFeed":
		n := i.Request.Intent.Slots.Name.Value
		q := "120"
		t := "formula"
		tm := "12"
		s := fmt.Sprintf("The last time %s ate, it was %s, and %s had %s of %s.", n, tm, n, q, t)
		resp.Say(s)
	case "GetNappyCount":
		n := i.Request.Intent.Slots.Name.Value
		from := "12"
		to := "1"
		q := "10"
		s := fmt.Sprintf("From %s to %s, %s had %s nappies changed.", from, to, n, q)
		resp.Say(s)
	case "GetFeedCount":
		n := i.Request.Intent.Slots.Name.Value
		from := "12"
		to := "1"
		q := "12"
		s := fmt.Sprintf("From %s to %s, %s ate %s times.", from, to, n, q)
		resp.Say(s)
	case "AMAZON.HelpIntent":
		resp.Say("This app is easy to use, just say: ask the office how warm it is")
	default:
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
