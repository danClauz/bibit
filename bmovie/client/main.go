package main

import (
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:1324", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect", err)
	}

	c := searchpb.NewSearchServiceClient(cc)
}
