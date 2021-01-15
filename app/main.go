package main

import (
	"context"
	"log"

	_ "github.com/tullo/microservice/client/haberdasher"
	"github.com/tullo/microservice/inject"
	pb "github.com/tullo/microservice/rpc/haberdasher"
)

func main() {
	// hd := haberdasher.New()
	// hd := haberdasher.NewCustom("http://172.28.0.4:3000", inject.NewHTTPClient())
	hd := pb.NewHaberdasherServiceProtobufClient("http://172.28.0.4:3000", inject.NewHTTPClient())
	hat, err := hd.MakeHat(context.Background(), &pb.Size{Centimeters: 58})
	if err != nil {
		panic(err)
	}
	log.Println(hat)
}
