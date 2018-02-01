package main

import (
    "fmt"
    proto "github.com/golang/protobuf/proto"
    "github.com/google/gtfs-realtime-bindings/golang/gtfs"
    "io/ioutil"
    "log"
    "net/http"
	"os"
)

type Credentials struct {
	Key string `json:"key"`
}

type line struct {
	url string
	color string
}

// Init map of all lines
var lines = map[string]*line{}

func init() {
	//Import API Key from ENV var MTAKEY
	var c Credentials
	c.Key = os.Getenv("MTAKEY")
	if len(c.Key) < 1 {
		log.Fatal("Error: Env var MTAKEY must contain a valid MTA API Key")
	}

	// Define vars
	ace_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=26"
	irt_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=1"
	nqrw_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=16"
	bdfm_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=21"
	l_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=2"
	g_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=31"
	jz_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=36"
	seven_url := "http://datamine.mta.info/mta_esi.php?key="+ c.Key + "&feed_id=51"

	for _, c := range "ace" {
		lines[string(c)] = &line{ace_url, "blue"}
	}
	for _, c := range "123" {
		lines[string(c)] = &line{irt_url, "red"}
	}
	for _, c := range "456" {
		lines[string(c)] = &line{irt_url, "green"}
	}

	for _, c := range "bdfm" {
		lines[string(c)] = &line{bdfm_url, "orange"}
	}
	for _, c := range "nqrw" {
		lines[string(c)] = &line{nqrw_url, "yellow"}
	}
	for _, c := range "jz" {
		lines[string(c)] = &line{jz_url, "brown"}
	}
	lines["s"] = &line{irt_url, "grey"}
	lines["l"] = &line{l_url, "grey"}
	lines["g"] = &line{g_url, "green"}
	lines["7"] = &line{seven_url, "purple"}
	//fmt.Printf("%v\n", lines["q"])

}

func main() {

	// Create our http client and pull the feed
    client := &http.Client{}
	req, err := http.NewRequest("GET", lines["a"].url, nil)
    //req.SetBasicAuth(username, password)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    feed := gtfs.FeedMessage{}
    err = proto.Unmarshal(body, &feed)
    if err != nil {
        log.Fatal(err)
    }
    for _, entity := range feed.Entity {
        if entity.TripUpdate != nil {
			tripUpdate := entity.TripUpdate
			trip := tripUpdate.Trip
			tripId := trip.TripId
			fmt.Printf("Trip ID: %s\n", *tripId)
		}
	}
}
