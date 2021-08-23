package server

import (
	"context"

	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svr *Server) HealthCheck(_ context.Context, _ *emptypb.Empty) (*searchpb.HealthCheckResponse, error) {
	return &searchpb.HealthCheckResponse{Status: "success"}, nil
}
