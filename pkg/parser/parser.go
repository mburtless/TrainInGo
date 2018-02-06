package parser

import (
    gtfs "github.com/mburtless/trainingo/pkg/transit_realtime"
    "bufio"
	"os"
	"log"
	//"fmt"
	"regexp"
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
	tripId := trip.GetTripId()
	r, _ := regexp.Compile(`\d{6}_\w\.{2}\w`)
	tripId = r.FindString(tripId)
	veh := Vehicle {
		Time:      vehPos.GetTimestamp(),
		Trip:      tripId,
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

func ParseStopSequences(filepath string, stops *map[string]*Stop) (*map[string]map[uint32]*Stop){
//func ParseStopSequences(filepath string, stops *map[string]*Stop) {
	// Takes filepath for stop_times.txt and pointer to stops map
	// Returns a map indexed on tripid containing slices of pointers to stops
	// Indexed on stop_sequence
	stopSequences := make(map[string]map[uint32]*Stop)

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal("Unable to read ", filepath, "! ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	r, _ := regexp.Compile(`\w{3}_\d{6}_\w+\.+\w`)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ",")
		tId := r.FindString(l[0])
		sId := l[3]
		//sSeq := strconv.FormatUint(l[4], 10)
		var sSeq uint32
		tmpSeq, err := strconv.ParseUint(l[4], 10, 32)
		if err != nil {
			//log.Fatal("Undable to parse stop lat/lon ", err)
			sSeq = uint32(0.0)
		} else {
			sSeq = uint32(tmpSeq)
		}
		//fmt.Printf("trip_id: %s station_id: %s stop_sequence: %s\n", tId, (*stops)[sId].StopName, sSeq)
		// check if tripid exists already, if so append to array of stop sequences
		if _, ok := stopSequences[tId]; ok {
			//s := &stopSequences[tId]
			//stopSequences[tId] = append(stopSequences[tId], (*stops)[sId])
			stopSequences[tId][sSeq] = (*stops)[sId]
		} else {
			//if not, create map, append this stop and add as elem in map
			//s := []*Stop{(*stops)[sId]}
			s := make(map[uint32]*Stop)
			s[sSeq] = (*stops)[sId]
			stopSequences[tId] = s
		}
	}
	//fmt.Printf("%v\n", stopSequences)
	return &stopSequences
}
func ParseStopId(veh *Vehicle) (string) {
	// Takes a vehicle and returns current stopid
	/*s := strings.Split(veh.Trip, "..")
	//dir := s[1]
	stopId := veh.Route + strconv.FormatUint(uint64(veh.StopSequence), 10) + dir*/
	stopId := strconv.FormatUint(uint64(veh.StopSequence), 10)
	if len(stopId) < 2 {
		stopId = "0" + stopId
	}

	// Search stop_times.txt for trip_id and stop seq
	// On match, save stop_id
	// Search stops.txt on stop_id
	// On match, save stop_name

	return stopId
}
