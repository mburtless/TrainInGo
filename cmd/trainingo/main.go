package main

import (
	"fmt"
    "github.com/mburtless/trainingo/pkg/feed"
    "github.com/mburtless/trainingo/pkg/parser"
    "github.com/mburtless/trainingo/configs"
	//"os"
    //"log"
	"time"
)

/*type Credentials struct {
	Key string `json:"key"`
}*/

/*type line struct {
	url string
	color string
}*/

// Init map of all lines
//var lines = map[string]*line{}

// Svc code required to determine weekend vs weekday trip ids
var svcCode string

func init() {
	//Import API Key from ENV var MTAKEY
	/*var c Credentials
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
	*/
	// Init svcCode based on today's date
	currentTime := time.Now()
	currentDay := currentTime.Weekday()
	switch currentDay {
		case 0:
			svcCode = "SUN"
		case 6:
			svcCode = "SAT"
		default:
			svcCode = "WKD"
	}
}

func main() {
	// Grab the api key and latest feed
	apiKey := configs.InitCredentials("MTAKEY")
	lineFeeds := configs.InitLineFeeds(apiKey)
	mtaFeed := *(feed.ReadFeed(lineFeeds["3"]))
	// Parse stops.txt for list of stops
	stops := *parser.ParseStops("third_party/nyct/stops.txt")
	// Parse stop_times.txt for list of trip ids correlated with their stop sequences and stops 
	stopSequences := *parser.ParseStopSequences("third_party/nyct/stop_times.txt", &stops)

	// Itterate through all VehiclePosition messages in the feed
	// Add them to slice of vehicles to track
    var vehicles []parser.Vehicle
	for _, entity := range mtaFeed.Entity {
		if entity.GetVehicle() != nil {
			// Probably a VehiclePosition message, parse it
			vehicles = append(vehicles, parser.ParseVehicle(entity))
		}
		/*if entity.TripUpdate != nil {
			tripUpdate := entity.TripUpdate
			trip := tripUpdate.Trip
			tripId := trip.TripId
			routeId := trip.RouteId
			fmt.Printf("Trip ID: %s\nRoute ID: %s\n\n", *tripId, *routeId)
		}*/
	}

	// Itterate through all vehicle positions found and print their current status
	for _, v := range vehicles {
		tId := svcCode + "_" + v.Trip
		vehStop := stopSequences[tId]
		if  v.StopSequence <= uint32(len(vehStop)) && v.StopSequence != 0 {
			fmt.Printf("%s is %s %s\n", tId, v.Status, vehStop[v.StopSequence].StopName)
		}
		// Handle the possibility that current stopsequence could be 0
		if v.StopSequence == 0 {
			fmt.Printf("Stopseq for %s is %d!\n", tId, v.StopSequence)
		}
	}
}
