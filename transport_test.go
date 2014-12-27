package transport

import (
	"testing"
	"log"
)

func TestGetClosestPlaces(t *testing.T) {
	types := []string{"train_station", "bus_station"}
	_, err := GetClosestPlaces("AIzaSyC50lfM-BNpgJMXesZ9qV4Jx6ubTMmwwxA", -33.859235, 151.068028, types)

	if err != nil {
		log.Fatal(err)
	}
}
