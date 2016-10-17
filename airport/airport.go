package airport

import (
	"math"

	"github.com/amy/project1_Protobuf/proto/airportdata"
)

type AirportServer interface {
	ClosestAirports()
}

type Airports struct {
	Airports []Airport
}

type Airport struct {
	Code      string
	Name      string
	Latitude  float64
	Longitude float64
	Distance  float64
}

type Request struct {
	Lat float64
	Lon float64
}

type Response struct {
	Airports [5]Airport
}

func dist(lat1, lon1, lat2, lon2 float64) float64 {

	rad := (math.Pi) / 180

	a := math.Sin(lat1*rad) * math.Sin(lat2*rad)
	b := math.Cos(lat1*rad) * math.Cos(lat2*rad)
	c := math.Cos(lat1*rad) * math.Sin(lat2*rad)

	return 111.12 * math.Acos(a+b*c)

}

func Store(list airportdata.AirportList) []Airport {

	a := []Airport{}

	for _, airport := range list.Airport {

		insert := Airport{
			Code:      *airport.Code,
			Name:      *airport.Name,
			Latitude:  *airport.Lat,
			Longitude: *airport.Lon,
			Distance:  0,
		}

		a = append(a, insert)
	}

	return a
}

func (a *Airports) ClosestAirports(args *Request, reply *Response) error {

	// current location's lat/lon
	lat := args.Lat
	lon := args.Lon

	// fill this up with closest 5
	var closest [5]Airport

	// range through the list of all airports
	// i is index in arrayList (aka slice)
	// airport is the airport value at index i
	for i, airport := range a.Airports {
		// add first 5 airports to closest
		airport.Distance = dist(lat, lon, airport.Latitude, airport.Longitude)
		if i < 5 {
			closest[i] = airport
			// past the first 5 we compare each airport to the farthest away airport in the list
			// if it's closer, we replace the farthest airport with the current one
		} else {

			largest := largestIndex(lat, lon, closest)

			// if current airport is closer than the farthest airport in "closest"
			// replace that airport with the current airport
			if airport.Distance < closest[largest].Distance {
				closest[largest] = airport
			}
		}
	}

	reply.Airports = closest

	return nil

}

// finds the index of the element in the "closest" array with largest distance
func largestIndex(lat1 float64, lon1 float64, array [5]Airport) int { // this might need double parentheses after Airport
	largest := 0
	for i := 1; i < 5; i++ {

		if array[i].Distance > array[largest].Distance {
			largest = i
		}
	}
	return largest
}
