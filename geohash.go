// The MIT License
//
// Copyright (c) 2012-2013 Maurício Magalhães
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package geolib

import "strings"

// GeoLocation holds geo location info
type GeoLocation struct {
	lat float64
	lon float64
}

// GeoEncode encodes GeoLocation
func GeoEncode(gl *GeoLocation, precision int) string {
	var geohash string
	var bit int
	var idx int32
	lonRange := []float64{-180, 180}
	latRange := []float64{-90, 90}
	even := true

	for len(geohash) < precision {
		lonMid := (lonRange[1] + lonRange[0]) / 2
		if even {
			if gl.lon > lonMid {
				idx |= bits[bit]
				lonRange[0] = lonMid
			} else {
				lonRange[1] = lonMid
			}
		} else {
			latMid := (latRange[1] + latRange[0]) / 2
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
	lonRange := []float64{-180, 180}
	latRange := []float64{-90, 90}

	for i := 0; i < l; i++ {
		c := string(hash[i])
		idx := strings.Index(geoBase32, c)

		for x := 4; x >= 0; x-- {
			bit := idx >> uint(x) & 1
			if even {
				lonMid := (lonRange[1] + lonRange[0]) / 2
				if bit == 1 {
					lonRange[0] = lonMid
				} else {
					lonRange[1] = lonMid
				}
			} else {
				latMid := (latRange[1] + latRange[0]) / 2
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

	return geo
}
