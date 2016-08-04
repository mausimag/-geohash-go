[![GoDoc](https://godoc.org/github.com/mausimag/geolib?status.svg)](https://godoc.org/github.com/mausimag/geolib)

Geo calculations
=======

Based on: https://en.wikipedia.org/wiki/Geohash

## Install

```console
go get github.com/mausimag/geolib
```

## Usage

```go
location := GeoLocation{
	lat: 48.668683,
	lon: -4.329321,
}

encoded := GeoEncode(&location, 9)
fmt.Printf("Encoded: [%s]", encoded)

decoded := GeoDecode("gbsuv7z7x")
fmt.Printf("Decoded: [%v, %v]", decoded.lat, decoded.lon)
```

## Methods

### GeoEncode(GeoLocation, precision) string

### GeoDecode(hashstring) GeoLocation

### DistanceBoundingCheck(sourceLat, sourceLon, distance) (topLeft GeoLocation, bottomRight GeoLocation)

### IsGeoLocationInArea(InitialGeoLocation, topLeft, bottomRight GeoLocation) bool
