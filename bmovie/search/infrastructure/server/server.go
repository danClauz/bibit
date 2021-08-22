package server

import (
	"context"
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
)

type Server struct {
	searchpb.UnimplementedSearchServiceServer
	sh shared.Holder
	ih interfaces.Holder
}

func New(sh shared.Holder, ih interfaces.Holder) *Server {
	return &Server{
		sh: sh,
		ih: ih,
	}
}

func (s *Server) SearchMovie(ctx context.Context, req *searchpb.SearchMovieRequest) (*searchpb.SearchMovieResponse, error) {
	resp, err := s.ih.SearchMovie(ctx, "", req.GetSearch(), int(req.GetPage()))
	return resp.SearchMovieResponseProto(), err
}
func (s *Server) DetailMovie(ctx context.Context, req *searchpb.DetailMovieRequest) (*searchpb.DetailMovieResponse, error) {
	resp, err := s.ih.DetailMovie(ctx, "", req.GetImdbId())
	return resp.DetailMovieResponseProto(), err
}
