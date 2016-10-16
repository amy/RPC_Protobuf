package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/amy/project1_Protobuf/place"
)

func main() {
	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &place.Request{
		Name:  "Rollins CDP",
		State: "MT",
	}
	var reply place.Response
	err = client.Call("Places.PlaceInfo", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Name: %v\n", args.Name)
	fmt.Printf("State: %v\n", args.State)
	fmt.Printf("Place: %v\n", reply)
}
