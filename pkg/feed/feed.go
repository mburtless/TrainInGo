package feed

import (
    proto "github.com/golang/protobuf/proto"
    "github.com/google/gtfs-realtime-bindings/golang/gtfs"
    "io/ioutil"
    "log"
	"net/http"
	"fmt"
)

func ReadFeed(url string) *gtfs.FeedMessage {
	// Takes a URL for the feed and returns a gtfs.FeedMessage{}

	// Create our http client and pull the feed
    client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
    //req.SetBasicAuth(username, password)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%p\n", &feed)
	return &feed
}
