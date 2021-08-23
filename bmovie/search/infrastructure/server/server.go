package server

import (
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
