package geolib

const (
	earthMinorAxis = 6356752.314245
	pi             = 3.14159265358979323846
	minLat         = -90.0
	maxLat         = 90.0
	minLon         = -180.0
	maxLon         = 180.0
	geoBase32      = "0123456789bcdefghjkmnpqrstuvwxyz"
)

var (
	bits = []int32{16, 8, 4, 2, 1}
)

func kmToM(kmv float64) float64 {
	return kmv * 1000
}

func toRadians(ang float64) float64 {
	return ang / 180.0 * pi
}

func toDegrees(angrad float64) float64 {
	return angrad * 180.0 / pi
}
