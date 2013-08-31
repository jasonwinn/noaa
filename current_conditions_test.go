package noaa

import (
	"fmt"
	"testing"
)

func TestCurrentConditions(t *testing.T) {
	seattleLat := 47.6064
	seattleLng := -122.330803
	nearestStation := "EBSW1"
	nearestStationLat := 47.605
	nearestStationLng := -122.338

	point := Point{Latitude: seattleLat, Longitude: seattleLng}
	station := point.NearestStation()
	c := station.CurrentConditions()

	if c.Latitude != nearestStationLat {
		t.Error(fmt.Sprintf("Expected %f, ~ Received %f", nearestStationLat, c.Latitude))
	}

	if c.Longitude != nearestStationLng {
		t.Error(fmt.Sprintf("Expected %f, ~ Received %f", nearestStationLng, c.Longitude))
	}

	if c.StationId != nearestStation {
		t.Error(fmt.Sprintf("Expected %s, ~ Received %s", nearestStation, c.StationId))
	}
}
