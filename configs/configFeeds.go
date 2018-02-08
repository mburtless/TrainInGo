package configs

import (
	"time"
)

func InitLineFeeds(apiKey string) (map[string]string) {
	// Initializes and returns our map of feeds
	var lineFeeds = map[string]string{}

	//Throw in something here to init credentials and set apiKey
	//apiKey := InitCredentials("MTAKEY")
	//Initialize our URLs based on our api key
	ace_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=26"
	irt_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=1"
	nqrw_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=16"
	bdfm_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=21"
	l_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=2"
	g_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=31"
	jz_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=36"
	seven_url := "http://datamine.mta.info/mta_esi.php?key="+ apiKey + "&feed_id=51"

	for _, c := range "ace" {
		lineFeeds[string(c)] = ace_url
	}
	for _, c := range "123" {
		lineFeeds[string(c)] = irt_url
	}
	for _, c := range "456" {
		lineFeeds[string(c)] = irt_url
	}

	for _, c := range "bdfm" {
		lineFeeds[string(c)] = bdfm_url
	}
	for _, c := range "nqrw" {
		lineFeeds[string(c)] = nqrw_url
	}
	for _, c := range "jz" {
		lineFeeds[string(c)] = jz_url
	}
	lineFeeds["s"] = irt_url
	lineFeeds["l"] = l_url
	lineFeeds["g"] = g_url
	lineFeeds["7"] = seven_url

	return lineFeeds
}

func InitSvcCode() (string) {
	var svcCode string
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

	return svcCode
}
