package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buddhamagnet/spotted/api"
)

var (
	spots map[string]string
	data  api.Spot
)

func init() {
	spots = map[string]string{
		"anders": "0i06IkOv1uEr8iXd1gbGWNgNpVLtt3XnH",
	}
}

func main() {
	res, err := http.Get(api.Endpoint + spots["anders"])
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(contents))
	fmt.Printf("%s\n", data.Response.FeedResponse.Feeds.Feed.Name)
}
