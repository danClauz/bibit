package main

import (
	"context"
	"log"

	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func main() {
	defer glog.Flush()
	ctx := context.Background()

	cc, err := grpc.Dial("127.0.0.1:2435", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect", err)
	}
	defer cc.Close()

	c := searchpb.NewSearchServiceClient(cc)

	searchResp, err := c.SearchMovie(ctx, &searchpb.SearchMovieRequest{
		Search: "batman",
		Page:   1,
	})

	if err != nil {
		glog.Fatalln(err)
	}

	glog.Info(searchResp)

	detailResp, err := c.DetailMovie(ctx, &searchpb.DetailMovieRequest{
		ImdbId: searchResp.GetSearch()[0].ImdbId,
	})

	if err != nil {
		glog.Fatalln(err)
	}

	glog.Info(detailResp)
}
