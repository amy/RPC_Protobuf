package airport

import "math"

type AirportServer interface {
	ClosestAirports()
}

type Airports struct {
	Airports []Airport
}

type Airport struct {
	code      string
	name      string
	latitude  float64
	longitude float64
	distance  float64
}

type Request struct {
	lat float64
	lon float64
}

type Response struct {
	Airports []Airport
}

func dist(lat1, lon1, lat2, lon2 float64) float64 {

	rad := (math.Pi) / 180

	a := math.Sin(lat1*rad) * math.Sin(lat2*rad)
	b := math.Cos(lat1*rad) * math.Cos(lat2*rad)
	c := math.Cos(lat1*rad) * math.Sin(lat2*rad)

	return 111.12 * math.Acos(a+b*c)

}

func (a *Airports) ClosestAirports(args *Request, reply *Response) error {

	// fill this up with closest 5
	closest := []Airport{}

	// range through the list of all airports
	// i is index in arrayList (aka slice)
	// airport is the airport value at index i
	for i, airport := range a.Airports {

	}

	return nil

}
