package main

import (
    "github.com/mburtless/trainingo/pkg/feed"
    "github.com/mburtless/trainingo/pkg/parser"
    "github.com/mburtless/trainingo/configs"
    "github.com/mburtless/trainingo/pkg/ui"
	//"time"
	//"flag"
)

// Svc code required to determine weekend vs weekday trip ids
//var svcCode string

/*func init() {
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
}*/

func main() {
	// Grab the api key and latest feed
	apiKey := configs.InitCredentials("MTAKEY")
	lineFeeds := configs.InitLineFeeds(apiKey)
	mtaFeed := *(feed.ReadFeed(lineFeeds["3"]))
	// Parse stops.txt for list of stops
	stops := *parser.ParseStops("third_party/nyct/stops.txt")
	// Parse stop_times.txt for list of trip ids correlated with their stop sequences and stops 
	stopSequences := *parser.ParseStopSequences("third_party/nyct/stop_times.txt", &stops)
	svcCode := configs.InitSvcCode()
	// Itterate through all VehiclePosition messages in the feed
	// Add them to slice of vehicles to track
    var vehicles []parser.Vehicle
	for _, entity := range mtaFeed.Entity {
		if entity.GetVehicle() != nil {
			// Probably a VehiclePosition message, parse it
			vehicles = append(vehicles, parser.ParseVehicle(entity))
		}
	}

	// Itterate through all vehicle positions found and print their current status
	/*for _, v := range vehicles {
		tId := svcCode + "_" + v.Trip
		vehStop := stopSequences[tId]
		if  v.StopSequence <= uint32(len(vehStop)) && v.StopSequence != 0 {
			fmt.Printf("%s is %s %s\n", tId, v.Status, vehStop[v.StopSequence].StopName)
		}
		// Handle the possibility that current stopsequence could be 0
		if v.StopSequence == 0 {
			fmt.Printf("Stopseq for %s is %d!\n", tId, v.StopSequence)
		}
	}*/
	ui.PrintVehiclePos(&vehicles, stopSequences, svcCode)
}
