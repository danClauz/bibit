package omdb

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/danClauz/bibit/bmovie/search/shared/config"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_client_SearchMovie(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	tests := []struct {
		name           string
		reqId          string
		searchKey      string
		page           int
		mockResponse   string
		mockStatusCode int
		want           *SearchResponse
		wantErr        bool
	}{
		{
			name:           "success",
			reqId:          "12345",
			searchKey:      "batman",
			page:           1,
			mockResponse:   mockSearchMovieResp,
			mockStatusCode: http.StatusOK,
			want: &SearchResponse{
				Search: []*search{
					{
						Title:  "Batman Begins",
						Year:   "2005",
						ImdbID: "tt0372784",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
					},
				},
				TotalResults: "445",
				Response:     "True",
			},
			wantErr: false,
		},
		{
			name:           "http status error",
			reqId:          "12345",
			searchKey:      "batman",
			page:           1,
			mockStatusCode: http.StatusBadRequest,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "unexpected request body",
			reqId:          "12345",
			searchKey:      "batman",
			page:           1,
			mockResponse:   `response body format changed`,
			mockStatusCode: http.StatusOK,
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))

			client := NewClient(shared.Holder{
				Logger: logger,
				Config: &config.EnvConfig{
					Omdb: &config.Omdb{
						Host: server.URL,
						Key:  "somekey",
					},
				},
			})
			assert.NotNil(client)

			output, err := client.SearchMovie(context.Background(), tt.reqId, tt.searchKey, tt.page)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, output)
		})
	}
}

func Test_client_SearchMovieByTitle(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	tests := []struct {
		name           string
		reqId          string
		searchKey      string
		mockResponse   string
		mockStatusCode int
		want           *SearchByTitleResponse
		wantErr        bool
	}{
		{
			name:           "success",
			reqId:          "12345",
			searchKey:      "batman",
			mockResponse:   mockSearchMovieByTitleResp,
			mockStatusCode: http.StatusOK,
			want: &SearchByTitleResponse{
				Title:    "Batman",
				Year:     "1989",
				Rated:    "PG-13",
				Released: "23 Jun 1989",
				Runtime:  "126 min",
				Genre:    "Action, Adventure",
				Director: "Tim Burton",
				Writer:   "Bob Kane, Sam Hamm, Warren Skaaren",
				Actors:   "Michael Keaton, Jack Nicholson, Kim Basinger",
				Plot:     "The Dark Knight of Gotham City begins his war on crime with his first major enemy being Jack Napier, a criminal who becomes the clownishly homicidal Joker.",
				Language: "English, French, Spanish",
				Country:  "United States, United Kingdom",
				Awards:   "Won 1 Oscar. 9 wins & 26 nominations total",
				Poster:   "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg",
				Ratings: []*rating{
					{
						Source: "Internet Movie Database",
						Value:  "7.5/10",
					},
					{
						Source: "Rotten Tomatoes",
						Value:  "71%",
					},
					{
						Source: "Metacritic",
						Value:  "69/100",
					},
				},
				Metascore:  "69",
				ImdbRating: "7.5",
				ImdbVotes:  "350,141",
				ImdbID:     "tt0096895",
				Type:       "movie",
				DVD:        "24 Jul 2014",
				BoxOffice:  "$251,348,343",
				Production: "Warner Brothers, Guber-Peters Company, PolyGram Filmed Entertainment",
				Website:    "N/A",
				Response:   "True",
			},
			wantErr: false,
		},
		{
			name:           "http status error",
			reqId:          "12345",
			searchKey:      "batman",
			mockStatusCode: http.StatusBadRequest,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "unexpected request body",
			reqId:          "12345",
			searchKey:      "batman",
			mockResponse:   `response body format changed`,
			mockStatusCode: http.StatusOK,
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))

			client := NewClient(shared.Holder{
				Logger: logger,
				Config: &config.EnvConfig{
					Omdb: &config.Omdb{
						Host: server.URL,
						Key:  "somekey",
					},
				},
			})
			assert.NotNil(client)

			output, err := client.SearchMovieByImdbId(context.Background(), tt.reqId, tt.searchKey)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, output)
		})
	}
}

const (
	mockSearchMovieResp = `{
"Search": [
{
"Title": "Batman Begins",
"Year": "2005",
"imdbID": "tt0372784",
"Type": "movie",
"Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
}
],
"totalResults": "445",
"Response": "True"
}`
	mockSearchMovieByTitleResp = `{
"Title": "Batman",
"Year": "1989",
"Rated": "PG-13",
"Released": "23 Jun 1989",
"Runtime": "126 min",
"Genre": "Action, Adventure",
"Director": "Tim Burton",
"Writer": "Bob Kane, Sam Hamm, Warren Skaaren",
"Actors": "Michael Keaton, Jack Nicholson, Kim Basinger",
"Plot": "The Dark Knight of Gotham City begins his war on crime with his first major enemy being Jack Napier, a criminal who becomes the clownishly homicidal Joker.",
"Language": "English, French, Spanish",
"Country": "United States, United Kingdom",
"Awards": "Won 1 Oscar. 9 wins & 26 nominations total",
"Poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg",
"Ratings": [
{
"Source": "Internet Movie Database",
"Value": "7.5/10"
},
{
"Source": "Rotten Tomatoes",
"Value": "71%"
},
{
"Source": "Metacritic",
"Value": "69/100"
}
],
"Metascore": "69",
"imdbRating": "7.5",
"imdbVotes": "350,141",
"imdbID": "tt0096895",
"Type": "movie",
"DVD": "24 Jul 2014",
"BoxOffice": "$251,348,343",
"Production": "Warner Brothers, Guber-Peters Company, PolyGram Filmed Entertainment",
"Website": "N/A",
"Response": "True"
}`
)
