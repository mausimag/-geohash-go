package geolib

import "math"

func DistanceBoundingCheck(slat, slon float64, distance float64) (topLeft GeoLocation, bottomRight GeoLocation) {
	radDist := kmToM(distance) / earthMinorAxis

	radLat := toRadians(slat)
	radLon := toRadians(slon)

	_minLat := radLat - radDist
	_maxLat := radLat + radDist

	var _minLon, _maxLon float64

	if _minLat > minLat && _maxLat < maxLat {
		deltaLon := math.Asin(math.Sin(radDist) / math.Cos(radLat))
		_minLon = radLon - deltaLon
		if _minLon < minLon {
			_minLon += 2 * pi
		}
		_maxLon = radLon + deltaLon
		if _maxLon > maxLon {
			_maxLon += 2 * pi
		}
	} else {
		_minLat = math.Max(_minLat, minLat)
		_maxLat = math.Min(_maxLat, maxLat)
		_minLon = minLon
		_maxLon = maxLon
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
