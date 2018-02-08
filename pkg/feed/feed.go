package feed

import (
    proto "github.com/golang/protobuf/proto"
    gtfs "github.com/mburtless/trainingo/pkg/transit_realtime"
    "io/ioutil"
    "log"
	"net/http"
	"time"
)

func ReadFeed(url string) *gtfs.FeedMessage {
	// Takes a URL for the feed and returns a gtfs.FeedMessage{}

	// Create our http client and pull the feed
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
    resp, err := client.Do(req)
    if err != nil {
		log.Fatalf("Error while requesting feed: ", err)
	} else if resp.Body == nil {
		log.Fatalf("Recieved empty response from %q", url)
	}
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatalf("Error while reading feed: ", err)
    }

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
    if err != nil {
		log.Fatalf("Error while unmarshaling feed: ", err)
    }

	return &feed
}
