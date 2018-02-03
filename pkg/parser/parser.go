package parser

import (
    gtfs "github.com/mburtless/trainingo/pkg/transit_realtime"
    //"log"
	//"fmt"
)


type Vehicle struct {
	Time           uint64
	Trip           string
	Route          string
	Status         gtfs.VehiclePosition_VehicleStopStatus
	StopSequence   uint32
}

func ParseVehicle(entity *gtfs.FeedEntity) (Vehicle) {
	// Takes an entity, if its a vehicle message parses and returns struct
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
