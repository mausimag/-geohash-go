package geolib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Geohash(t *testing.T) {
	assert := assert.New(t)

	location := GeoLocation{
		lat: 48.668683,
		lon: -4.329321,
	}

	encoded := GeoEncode(&location, 9)

	t.Logf("Encoded: [%s]", encoded)

	decoded := GeoDecode("gbsuv7z7x")

	t.Logf("Decoded: [%v, %v]", decoded.lat, decoded.lon)

	assert.Equal(location.lat, location.lat, "Wrong lat")
	assert.Equal(decoded.lon, decoded.lon, "Wrong lon")
}

func Test_DistanceBoundingCheck(t *testing.T) {
	assert := assert.New(t)

	var points = []GeoLocation{
		GeoLocation{
			lat: -20.279877,
			lon: 57.518932,
		},
		GeoLocation{
			lat: -20.279343,
			lon: 57.518374,
		},
		GeoLocation{
			lat: -20.282030,
			lon: 57.520021,
		},
		GeoLocation{
			lat: -20.290418,
			lon: 57.535239,
		},
	}

	topLeft, bottomRight := DistanceBoundingCheck(points[0].lat, points[0].lon, 1.0)
	t.Logf("Result: topLeft: [%v, %v], bottomRight: [%v, %v]", topLeft.lat, topLeft.lon, bottomRight.lat, bottomRight.lon)

	expected := []bool{true, true, true, false}

	for idx, p := range points {
		r := IsGeoLocationInArea(&p, &topLeft, &bottomRight)
		t.Logf("Decoded: [%v, %v] - %v", p.lat, p.lon, r)
		assert.Equal(expected[idx], r)
	}
}
