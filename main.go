package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/buddhamagnet/spotted/api"
	"github.com/mgutz/ansi"
)

var (
	spots map[string]string
)

func init() {
	spots = map[string]string{
		"anders": "0i06IkOv1uEr8iXd1gbGWNgNpVLtt3XnH",
	}
}

var (
	phosphorize = ansi.ColorFunc("green+h:black")
	errorize    = ansi.ColorFunc("red+h:black")
)

func main() {
	data := make(chan string, 10)
	go poll(data)
	go receive(data)
	fmt.Scanf("%s", os.Stdout)
}

func poll(data chan string) {
	for {
		res, err := http.Get(api.Endpoint + spots["anders"])
		if err != nil {
			data <- errorize(fmt.Sprintf("%v", err))
			continue
		}
		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
			data <- errorize(fmt.Sprintf("%v", err))
			continue
		}
		var spot api.Spot
		if err := json.Unmarshal(contents, &spot); err != nil {
			data <- errorize(fmt.Sprintf("%v", err))
			continue
		}
		data <- fmt.Sprintf("%s\n", string(contents))
		data <- phosphorize(fmt.Sprintf("%s: %s\n", time.Now().UTC(), spot.Response.FeedResponse.Feeds.Feed.Name))
		time.Sleep(5 * time.Second)
	}
}

func receive(data chan string) {
	for {
		response := <-data
		fmt.Println(response)
	}
}
