package cbutil

import (
	"math"
	"strconv"
	"strings"
)

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func DistanceBetweenLatLongs(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func GetLatLongFromString(latLong string) (float64, float64) {
	if len(latLong) == 0 || !strings.Contains(latLong, ",") {
		latLong = "40.236422,-100.094776"
	}

	split := strings.Split(latLong, ",")

	split[0] = strings.TrimSpace(split[0])
	split[1] = strings.TrimSpace(split[1])

	var lat, long float64

	lat, e := strconv.ParseFloat(split[0], 32)
	if e != nil {
		lat = 40.236422
	}

	long, e = strconv.ParseFloat(split[1], 32)
	if e != nil {
		long = -100.094776
	}

	return lat, long
}