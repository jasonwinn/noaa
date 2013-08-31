package noaa

import (
	"fmt"
	"testing"
)

func TestPointDistance(t *testing.T) {
	seattleLat := 47.6064
	seattleLng := -122.330803
	vancouverLat := 49.261226
	vancouverLng := -123.113927
	expectedDistance := 192.86249660526974

	seattle := &Point{
		Latitude:  seattleLat,
		Longitude: seattleLng,
	}

	vancouver := &Point{
		Latitude:  vancouverLat,
		Longitude: vancouverLng,
	}

	distance := seattle.HaversineDistance(vancouver)

	if distance != expectedDistance {
		t.Error(fmt.Sprintf("Expected %f, ~ Received %f", expectedDistance, distance))
	}
}
