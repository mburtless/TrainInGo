package parser

import (
    gtfs "github.com/mburtless/trainingo/pkg/transit_realtime"
    //"log"
	"fmt"
)


type Vehicle struct {
	Time      uint64
	Trip      string
	Route     string
	Status    string
	StopSeq   string
}

func ParseVehicle(entity *gtfs.FeedEntity) {
	// Takes an entity, if its a vehicle message parses and returns struct
	var vehPos *gtfs.VehiclePosition = entity.GetVehicle()
	if vehPos != nil {
		fmt.Printf("Vehicle: %v", vehPos)
	}
}
