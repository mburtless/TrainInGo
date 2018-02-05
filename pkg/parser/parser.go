package parser

import (
    gtfs "github.com/mburtless/trainingo/pkg/transit_realtime"
    "bufio"
	"os"
	"log"
	//"fmt"
	"strings"
	"strconv"
)


type Vehicle struct {
	Time           uint64
	Trip           string
	Route          string
	Status         gtfs.VehiclePosition_VehicleStopStatus
	StopSequence   uint32
}

type Stop struct {
	StopId        string
	StopName      string
	StopLat       uint64
	StopLon       uint64
}

func ParseVehicle(entity *gtfs.FeedEntity) (Vehicle) {
	// Takes a vehicle message entity parses and returns struct
	var vehPos *gtfs.VehiclePosition = entity.GetVehicle()
	var trip *gtfs.TripDescriptor = vehPos.GetTrip()
	veh := Vehicle {
		Time:      vehPos.GetTimestamp(),
		Trip:      trip.GetTripId(),
		Route:     trip.GetRouteId(),
		Status:    vehPos.GetCurrentStatus(),
		StopSequence:   vehPos.GetCurrentStopSequence(),
	}

	return veh
}

func ParseStops(filepath string) (*map[string]*Stop){
	// Takes filepath for stops.txt and 
	// parses it to map of Stop stucts
	stops := make(map[string]*Stop)

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal("Unable to read ", filepath, "! ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ",")
		sLat, err := strconv.ParseUint(l[4], 10, 64)
		sLon, err := strconv.ParseUint(l[5], 10, 64)
		if err != nil {
			//log.Fatal("Undable to parse stop lat/lon ", err)
			sLat, sLon = 0.0, 0.0
		}
		s := Stop {
			StopId:		l[0],
			StopName:	l[2],
			StopLat:	sLat,
			StopLon:	sLon,
		}
		stops[l[0]] = &s
		//fmt.Println(scanner.Text())
	}

	return &stops
}
