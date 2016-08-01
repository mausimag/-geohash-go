package geohash

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	encoded := GeoEncode(&GeoLocation{
		lat:       48.668683,
		lon:       -4.329321,
		precision: 9,
	})

	log.Printf("Encoded: %s", encoded)

	decoded := GeoDecode("gbsuv7z7x")

	log.Printf("Decoded: [%v, %v] - %v", decoded.lat, decoded.lon, decoded.precision)
}
