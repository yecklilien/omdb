package impl

import (
	"net/http"

	"github.com/yecklilien/OMDB/movie"
)

//MovieAPIHTTPJSONRPCServer to be implemented
type MovieAPIHTTPJSONRPCServer struct {
	movieAPI *MovieAPI
}

//NewMovieAPIHTTPJSONRPCServer function return new instance of MovieAPIHTTPJSONRPCServer
func NewMovieAPIHTTPJSONRPCServer(movieAPI *MovieAPI) *MovieAPIHTTPJSONRPCServer {
	return &MovieAPIHTTPJSONRPCServer{
		movieAPI: movieAPI,
	}
}

//SearchMovie http json rpc implementation
func (server *MovieAPIHTTPJSONRPCServer) SearchMovie(r *http.Request, spec *movie.SearchMovieRequest, result *movie.SearchMovieResponse) error {
	searchResult, err := server.movieAPI.SearchMovie(spec)
	*result = movie.SearchMovieResponse{
		Movies:      searchResult.Movies,
		TotalResult: searchResult.TotalResult,
	}
	return err
}

//GetMovieDetail http json rpc implementation
func (server *MovieAPIHTTPJSONRPCServer) GetMovieDetail(r *http.Request, spec *movie.GetMovieDetailRequest, result *movie.GetMovieDetailResponse) error {
	getMovieDetailResult, err := server.movieAPI.GetMovieDetail(spec)
	*result = movie.GetMovieDetailResponse{
		Title:      getMovieDetailResult.Title,
		Year:       getMovieDetailResult.Year,
		Rated:      getMovieDetailResult.Rated,
		Released:   getMovieDetailResult.Released,
		Runtime:    getMovieDetailResult.Runtime,
		Genre:      getMovieDetailResult.Genre,
		Director:   getMovieDetailResult.Director,
		Writer:     getMovieDetailResult.Writer,
		Actors:     getMovieDetailResult.Actors,
		Plot:       getMovieDetailResult.Plot,
		Language:   getMovieDetailResult.Language,
		Country:    getMovieDetailResult.Country,
		Awards:     getMovieDetailResult.Awards,
		Poster:     getMovieDetailResult.Poster,
		Ratings:    getMovieDetailResult.Ratings,
		Metascore:  getMovieDetailResult.Metascore,
		ImdbRating: getMovieDetailResult.ImdbRating,
		ImdbVotes:  getMovieDetailResult.ImdbVotes,
		ImdbID:     getMovieDetailResult.ImdbID,
		Type:       getMovieDetailResult.Type,
		DVD:        getMovieDetailResult.DVD,
		BoxOffice:  getMovieDetailResult.BoxOffice,
		Production: getMovieDetailResult.Production,
		Website:    getMovieDetailResult.Website,
	}
	return err
}
