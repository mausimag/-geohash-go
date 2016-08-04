package geolib

const (
	geoEarthMinorAxis = 6356752.314245
	geoPI             = 3.14159265358979323846
	geoMinLat         = -90.0
	geoMaxLat         = 90.0
	geoMinLon         = -180.0
	geoMaxLon         = 180.0
	geoBase32      = "0123456789bcdefghjkmnpqrstuvwxyz"
)

var (
	bits = []int32{16, 8, 4, 2, 1}
)

func kmToM(kmv float64) float64 {
	return kmv * 1000
}

func toRadians(ang float64) float64 {
	return ang / 180.0 * geoPI
}

func toDegrees(angrad float64) float64 {
	return angrad * 180.0 / geoPI
}
