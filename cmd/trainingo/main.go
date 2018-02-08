package main

import (
    "github.com/mburtless/trainingo/pkg/feed"
    "github.com/mburtless/trainingo/pkg/parser"
    "github.com/mburtless/trainingo/configs"
    "github.com/mburtless/trainingo/pkg/ui"
	"flag"
	"log"
	"os"
)

func verifyLine(l *string) (bool) {
	lineChoices := map[string]bool{}
	for _, c := range "acenqrwbdfmlgjzs1234567" {
		lineChoices[string(c)] = true
	}

	if _, validChoice := lineChoices[*l]; !validChoice {
		log.Fatalf("-line provided %q is not a valid MTA subway line", *l)
	}

	return true
}

func main() {
	// Init CLI args
	linePtr := flag.String("line", "", "Subway line to fetch data on (Required)")
	flag.Parse()

	// Make sure line arg was provided and is valid
	if *linePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		verifyLine(linePtr)
	}


	// Grab the api key env, svc code and latest feed
	apiKey := configs.InitCredentials("MTAKEY")
	lineFeeds := configs.InitLineFeeds(apiKey)
	mtaFeed := *(feed.ReadFeed(lineFeeds[*linePtr]))
	svcCode := configs.InitSvcCode()

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
	}

	ui.PrintVehiclePos(&vehicles, stopSequences, svcCode)
}
