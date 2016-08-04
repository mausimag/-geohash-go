package geolib

import "math"

func DistanceBoundingCheck(slat, slon float64, distance float64) (topLeft GeoLocation, bottomRight GeoLocation) {
	radDist := kmToM(distance) / geoEarthMinorAxis

	radLat := toRadians(slat)
	radLon := toRadians(slon)

	_minLat := radLat - radDist
	_maxLat := radLat + radDist

	var _minLon, _maxLon float64

	if _minLat > geoMinLat && _maxLat < geoMaxLat {
		deltaLon := math.Asin(math.Sin(radDist) / math.Cos(radLat))
		_minLon = radLon - deltaLon
		if _minLon < geoMinLon {
			_minLon += 2 * geoPI
		}
		_maxLon = radLon + deltaLon
		if _maxLon > geoMaxLon {
			_maxLon += 2 * geoPI
		}
	} else {
		_minLat = math.Max(_minLat, geoMinLat)
		_maxLat = math.Min(_maxLat, geoMaxLat)
		_minLon = geoMinLon
		_maxLon = geoMaxLon
	}

	topLeft = GeoLocation{
		lat: toDegrees(_maxLat),
		lon: toDegrees(_minLon),
	}

	bottomRight = GeoLocation{
		lat: toDegrees(_minLat),
		lon: toDegrees(_maxLon),
	}

	return
}

func IsGeoLocationInArea(gl *GeoLocation, lTopLeft, lBottomRight *GeoLocation) bool {
	return (gl.lat < lTopLeft.lat && gl.lat > lBottomRight.lat) && (gl.lon > lTopLeft.lon && gl.lon < lBottomRight.lon)
}
