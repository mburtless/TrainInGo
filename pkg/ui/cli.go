package ui

import (
	"fmt"
	"github.com/mburtless/trainingo/pkg/parser"
)

func PrintVehiclePos(vehicles *[]parser.Vehicle, stopSequences map[string]map[uint32]*parser.Stop, svcCode string) {
	for _, v := range *(vehicles) {
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
