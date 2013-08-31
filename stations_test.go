package noaa

import (
	"fmt"
	"testing"
)

func TestNearestStation(t *testing.T) {
	seattleLat := 47.6064
	seattleLng := -122.330803
	nearestStation := "EBSW1"
	nearestStationLat := 47.605
	nearestStationLng := -122.338

	point := Point{Latitude: seattleLat, Longitude: seattleLng}
	station := point.NearestStation()

	if station.Latitude != nearestStationLat {
		t.Error(fmt.Sprintf("Expected %f, ~ Received %f", nearestStationLat, station.Latitude))
	}

	if station.Longitude != nearestStationLng {
		t.Error(fmt.Sprintf("Expected %f, ~ Received %f", nearestStationLng, station.Longitude))
	}

	if station.Id != nearestStation {
		t.Error(fmt.Sprintf("Expected %s, ~ Received %s", nearestStation, station.Id))
	}
}
