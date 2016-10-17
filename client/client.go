package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/amy/project1_Protobuf/airport"
	"github.com/amy/project1_Protobuf/place"
)

func main() {

	city := "Rollins CDP"
	state := "MT"

	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1235")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Call placeServer to look up place long & lat
	args := place.Request{
		Name:  city,
		State: state,
	}
	var reply place.Response
	err = client.Call("Places.PlaceInfo", &args, &reply)
	if err != nil {
		log.Fatal("places error:", err)
	}

	longitude := reply.Place.Lon
	latitude := reply.Place.Lat

	fmt.Println(longitude)
	fmt.Println(latitude)

	// Call airportServer to look up closest 5 airports
	serverAddress1 := "localhost"
	client1, err := rpc.DialHTTP("tcp", serverAddress1+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args1 := airport.Request{
		Lat: latitude,
		Lon: longitude,
	}

	var reply1 airport.Response
	err = client1.Call("Airports.ClosestAirports", &args1, &reply1)
	if err != nil {
		log.Fatal("airport error:", err)
	}

	fmt.Println(reply1)
}
