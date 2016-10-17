package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/amy/project1_Protobuf/place"
	"github.com/amy/project1_Protobuf/proto/placedata"
	"github.com/golang/protobuf/proto"
)

func main() {

	//////////////
	// PROTOBUF //
	//////////////

	data, err := ioutil.ReadFile("../proto/placedata/places-proto.bin")
	if err != nil {
		log.Fatal(err)
	}

	var list placedata.PlaceList
	err = proto.Unmarshal(data, &list)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	////////////
	// SERVER //
	////////////

	places := place.Places{}
	places.Store = place.Store(list)

	rpc.Register(&places)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1235")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
