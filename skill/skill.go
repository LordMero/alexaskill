package main

import (
	"alexaskill/skill/models"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func HandleRequest(ctx context.Context, i models.AlexaComplexRequest) (models.AlexaResponse, error) {
	// Use Spew to output the request for debugging purposes:
	fmt.Println("---- Dumping Input Map: ----")
	spew.Dump(i)
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	log.Printf("Request type is ", i.Request.Intent.Name)

	// Create a response object
	resp := models.CreateResponse()

	// Customize the response for each Alexa Intent
	switch i.Request.Intent.Name {
	case "AddNappy":
		t := i.Request.Intent.Slots.Kind.Value
		s := fmt.Sprintf("I am adding a %s nappy. I hope it doesn't smell!", t)
		resp.Say(s)
	case "AddFeed":
		t := i.Request.Intent.Slots.Type.Value
		//q, _ := strconv.ParseFloat(i.Request.Intent.Slots.Quantity.Value, 64)
		q := i.Request.Intent.Slots.Quantity.Value
		s := fmt.Sprintf("I am adding %s of %s. Thanks", q, t)
		resp.Say(s)
	case "AddBreastFeed":
		logType := i.Request.Intent.Slots.LogType.Value
		dur := i.Request.Intent.Slots.Duration.Value
		timeUnit := i.Request.Intent.Slots.TimeUnit.Value

		s := fmt.Sprintf("I am %sing a breast feeding session of %s %s.", logType, dur, timeUnit)
		resp.Say(s)
	case "AddWeight":
		n := i.Request.Intent.Slots.Name.Value
		q := i.Request.Intent.Slots.Wgt.Value
		s := fmt.Sprintf("%s now weights %s. Got it!", n, q)
		resp.Say(s)
	case "GetWeight":
		n := i.Request.Intent.Slots.Name.Value
		q := "4.4"
		s := fmt.Sprintf("%s weights %s kilos!", n, q)
		resp.Say(s)
	case "GetLastFeedTime":
		n := i.Request.Intent.Slots.Name.Value
		q := "12"
		s := fmt.Sprintf("%s lastly ate at %s", n, q)
		resp.Say(s)
	case "GetLastFeedQuantity":
		n := i.Request.Intent.Slots.Name.Value
		q := "12"
		t := "formula"
		s := fmt.Sprintf("Last feed %s had %s of %s", n, q, t)
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
