package main

import (
	"alexaskill/models"
	"alexaskill/utilities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://grazianomirata.com:8080/api/baby/")
	utilities.Catch(err)

	defer resp.Body.Close()


	utilities.Catch(err)

	body, err := ioutil.ReadAll(resp.Body)
	utilities.Catch(err)
	fmt.Println(body)

	w := models.Baby{}
	err = json.Unmarshal(body, &w)
	utilities.Catch(err)
	fmt.Println(w)



}
