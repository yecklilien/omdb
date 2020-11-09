package entity

//SearchMovieResult Entity
type SearchMovieResult struct {
	Movies []Movie `json:"Search"`
	TotalResult string `json:"totalResults"`
	Response string `json:"Response"`
	Error string `json:"Error"`	
}

//Movie Entity
type Movie struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}

//MovieDetail Entity
type MovieDetail struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"` 
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`  
	ImdbID     string `json:"imdbID"` 
	Type       string `json:"Type"` 
	DVD        string `json:"DVD"` 
	BoxOffice  string `json:"BoxOffice"` 
	Production string `json:"Production"` 
	Website    string `json:"Website"`
	Response   string `json:"Response"`
	Error      string `json:"Error"`	
} 

//Rating Entity
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

//SearchMovieSpec Entity
type SearchMovieSpec struct {
	Query string
	Page int32
} 

//GetMovieDetailSpec Entity
type GetMovieDetailSpec struct {
	ImdbID string
} 