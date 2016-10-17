package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/amy/project1_Protobuf/proto/airportdata"
	"github.com/golang/protobuf/proto"
)

func main() {

	//////////////
	// PROTOBUF //
	//////////////

	data, err := ioutil.ReadFile("../proto/airportdata/airports-proto.bin")
	if err != nil {
		log.Fatal(err)
	}

	var list airportdata.AirportList
	err = proto.Unmarshal(data, &list)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	////////////
	// SERVER //
	////////////

	rpc.Register(&airports)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

}
