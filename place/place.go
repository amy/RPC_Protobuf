package place

import (
	"fmt"
	"strings"

	"github.com/amy/project1_Protobuf/proto/placedata"
)

type PlaceService interface {
	PlaceInfo()
}

type Places struct {
	Store map[string][2]float64
}

type Place struct {
	State string
	Name  string
	Lat   float64
	Lon   float64
}

type Request struct {
	Name  string
	State string
}

type Response struct {
	Place Place
}

func (a *Places) PlaceInfo(args *Request, reply *Response) error {

	key := strings.ToUpper(args.State) + args.Name

	value := a.Store[key]

	reply.Place = Place{
		State: strings.ToUpper(args.State),
		Name:  args.Name,
		Lat:   value[0],
		Lon:   value[1],
	}

	fmt.Printf("Lat %v\n", value[0])

	return nil
}

func Store(list placedata.PlaceList) map[string][2]float64 {

	lookup := make(map[string][2]float64)

	for i, _ := range list.Place {
		p := list.Place[i]
		key := p.GetState() + p.GetName()
		value := [2]float64{*p.Lat, *p.Lon}
		lookup[key] = value
	}

	return lookup
}
