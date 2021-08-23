package server

import (
	"context"

	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
)

func (svr *Server) SearchMovie(ctx context.Context, req *searchpb.SearchMovieRequest) (*searchpb.SearchMovieResponse, error) {
	resp, err := svr.ih.SearchMovie(ctx, "", req.GetSearch(), int(req.GetPage()))
	if err != nil {
		return nil, err
	}
	return resp.SearchMovieResponseProto(), nil
}

func (svr *Server) DetailMovie(ctx context.Context, req *searchpb.DetailMovieRequest) (*searchpb.DetailMovieResponse, error) {
	resp, err := svr.ih.DetailMovie(ctx, "", req.GetImdbId())
	if err != nil {
		return nil, err
	}
	return resp.DetailMovieResponseProto(), nil
}
