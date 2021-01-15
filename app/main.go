package main

import (
	"context"
	"log"

	"github.com/tullo/microservice/client/haberdasher"
	"github.com/tullo/microservice/inject"
	rp "github.com/tullo/microservice/rpc/haberdasher"
)

func main() {
	// hd := haberdasher.New()
	// hd := haberdasher.NewHaberdasherServiceProtobufClient("http://172.28.0.4:3000", inject.NewHTTPClient())
	hd := haberdasher.NewCustom("http://172.28.0.4:3000", inject.NewHTTPClient())
	hat, err := hd.MakeHat(context.Background(), &rp.Size{Centimeters: 78})
	if err != nil {
		panic(err)
	}
	log.Println(hat)
}
