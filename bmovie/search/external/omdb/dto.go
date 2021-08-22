package omdb

type (
	SearchResponse struct {
		Search       []*search `json:"Search"`
		TotalResults string    `json:"totalResults"`
		Response     string    `json:"Response"`
	}
	search struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}
	SearchByTitleResponse struct {
		Title      string  `json:"Title"`
		Year       string  `json:"Year"`
		Rated      string  `json:"Rated"`
		Released   string  `json:"Released"`
		Runtime    string  `json:"Runtime"`
		Genre      string  `json:"Genre"`
		Director   string  `json:"Director"`
		Writer     string  `json:"Writer"`
		Actors     string  `json:"Actors"`
		Plot       string  `json:"Plot"`
		Language   string  `json:"Language"`
		Country    string  `json:"Country"`
		Awards     string  `json:"Awards"`
		Poster     string  `json:"Poster"`
		Ratings    []*rating `json:"Ratings"`
		Metascore  string  `json:"Metascore"`
		ImdbRating string  `json:"imdbRating"`
		ImdbVotes  string  `json:"imdbVotes"`
		ImdbID     string  `json:"imdbID"`
		Type       string  `json:"Type"`
		DVD        string  `json:"DVD"`
		BoxOffice  string  `json:"BoxOffice"`
		Production string  `json:"Production"`
		Website    string  `json:"Website"`
		Response   string  `json:"Response"`
	}
	rating struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	}
)
