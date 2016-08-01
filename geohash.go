package geohash

import "strings"

// GeoLocation holds geo location info
type GeoLocation struct {
	lat       float32
	lon       float32
	precision uint16
}

var (
	latRange  = []float32{-90.0, 90.0}
	lonRange  = []float32{-180.0, 180.0}
	bits      = []int32{16, 8, 4, 2, 1}
	geoBase32 = "0123456789bcdefghjkmnpqrstuvwxyz"
)

func middle(r []float32) float32 {
	return (r[0] + r[1]) / 2
}

// GeoEncode encodes GeoLocation
func GeoEncode(gl *GeoLocation) string {
	var geohash string
	var bit int
	var idx int32
	even := true

	for len(geohash) < int(gl.precision) {
		lonMid := middle(lonRange)
		if even {
			if gl.lon > lonMid {
				idx |= bits[bit]
				lonRange[0] = lonMid
			} else {
				lonRange[1] = lonMid
			}
		} else {
			latMid := middle(latRange)
			if gl.lat > latMid {
				idx |= bits[bit]
				latRange[0] = latMid
			} else {
				latRange[1] = latMid
			}
		}

		if bit < 4 {
			bit++
		} else {
			geohash += string(geoBase32[idx])
			bit = 0
			idx = 0
		}

		even = !even
	}

	return geohash
}

// GeoDecode decodes hash
func GeoDecode(hash string) *GeoLocation {
	geo := &GeoLocation{}
	l := len(hash)
	even := true

	for i := 0; i < l; i++ {
		c := string(hash[i])
		idx := strings.Index(geoBase32, c)

		for x := 4; x >= 0; x-- {
			bit := idx >> uint(x) & 1
			if even {
				lonMid := (lonRange[0] + lonRange[1]) / 2
				if bit == 1 {
					lonRange[0] = lonMid
				} else {
					lonRange[1] = lonMid
				}
			} else {
				latMid := (latRange[0] + latRange[1]) / 2
				if bit == 1 {
					latRange[0] = latMid
				} else {
					latRange[1] = latMid
				}
			}

			even = !even
		}
	}

	geo.lat = (latRange[0] + latRange[1]) / 2
	geo.lon = (lonRange[0] + lonRange[1]) / 2
	geo.precision = uint16(l)

	return geo
}
