package model

import (
	"time"

	"github.com/danClauz/bibit/bmovie/search/external/omdb"
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
)

type SearchHistory struct {
	ID        uint
	RequestID string
	SearchKey string
	Page      int
	Result    interface{}
	CreatedAt time.Time
}

func (*SearchHistory) TableName() string {
	return "search_history"
}

func (s *SearchHistory) SearchMovieResponseProto() *searchpb.SearchMovieResponse {
	if res, ok := s.Result.(*omdb.SearchResponse); ok {
		search := make([]*searchpb.Search, 0)

		for _, val := range res.Search {
			search = append(search, &searchpb.Search{
				Title:  val.Title,
				Year:   val.Year,
				ImdbId: val.ImdbID,
				Type:   val.Type,
				Poster: val.Poster,
			})
		}

		return &searchpb.SearchMovieResponse{
			Search:       search,
			TotalResults: res.TotalResults,
			Response:     res.Response,
		}
	}

	return nil
}

func (s *SearchHistory) DetailMovieResponseProto() *searchpb.DetailMovieResponse {
	if res, ok := s.Result.(*omdb.SearchByTitleResponse); ok {
		ratings := make([]*searchpb.Rating, 0)

		for _, val := range res.Ratings {
			ratings = append(ratings, &searchpb.Rating{
				Source: val.Source,
				Value:  val.Value,
			})
		}

		return &searchpb.DetailMovieResponse{
			Title:      res.Title,
			Year:       res.Year,
			Rated:      res.Rated,
			Released:   res.Released,
			Runtime:    res.Runtime,
			Genre:      res.Genre,
			Director:   res.Director,
			Writer:     res.Writer,
			Actors:     res.Actors,
			Plot:       res.Plot,
			Language:   res.Language,
			Country:    res.Country,
			Awards:     res.Awards,
			Poster:     res.Poster,
			Ratings:    ratings,
			Metascore:  res.Metascore,
			ImdbRating: res.ImdbRating,
			ImdbVotes:  res.ImdbVotes,
			ImdbId:     res.ImdbID,
			Type:       res.Type,
			Dvd:        res.DVD,
			BoxOffice:  res.BoxOffice,
			Production: res.Production,
			Website:    res.Website,
			Response:   res.Response,
		}
	}

	return nil
}
