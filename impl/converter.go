package impl

import (
	"errors"
	"log"
	"github.com/yecklilien/OMDB/movie"
	"github.com/yecklilien/OMDB/logger/entity"
	"encoding/json"
	"strconv"
	"time"
)

type searchMovieResponseOMDB struct {
	Movies []movieOMDB `json:"Search"`
	TotalResult string `json:"totalResults"`
	Response string `json:"Response"`
	Error string `json:"Error"`	
}

type movieOMDB struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}

type movieDetailOMDB struct {
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
	Ratings    []ratingOMDB `json:"Ratings"`
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

type ratingOMDB struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}


func constructSearchMovieResponse(bytes []byte) (*movie.SearchMovieResponse, error) {
	var omdbResponse searchMovieResponseOMDB
	err:= json.Unmarshal(bytes,&omdbResponse)
	
	if err!=nil {
		log.Fatalln(err);
		return nil,err
	}

	if omdbResponse.Response != "True" {
		log.Fatalln(omdbResponse.Error);
		return nil, errors.New(omdbResponse.Error)
	}
	
	result,err := getSearchMovieResponse(omdbResponse)
	
	if err!=nil {
		log.Fatalln(err);
		return nil,err
	}

	return result,nil
}

func getSearchMovieResponse(spec searchMovieResponseOMDB) (*movie.SearchMovieResponse, error) {
	
	resultChan:= make(chan *movie.Movie)

	for _,v := range spec.Movies {
		go getMovie(v, resultChan)
	}
	
	movies := make([]*movie.Movie, len(spec.Movies))
	for i := range movies {
		movies[i] = <-resultChan
	}

	totalResults,err := strconv.Atoi(spec.TotalResult);

	if(err!=nil) {
		log.Fatalln(err)
		return nil,err;
	}

	searchMovieResponse := movie.SearchMovieResponse {
		Movies: movies,
		TotalResult: int32(totalResults),
	}

	return &searchMovieResponse,nil
}

func getMovie(spec movieOMDB, result chan *movie.Movie){
	result <- &movie.Movie {
		Title : spec.Title,
		Year : spec.Year,
		ImdbID : spec.ImdbID,
		Poster : spec.Poster,
		Type : spec.Type,
	}
}

func constructGetMovieDetailResponse(bytes []byte) (*movie.GetMovieDetailResponse, error) {
	var omdbResponse movieDetailOMDB
	err:= json.Unmarshal(bytes,&omdbResponse)
	
	if err!=nil {
		log.Fatalln(err);
		return nil,err
	}

	if omdbResponse.Response != "True" {
		log.Fatalln(omdbResponse.Error);
		return nil, errors.New(omdbResponse.Error)
	}

	return getGetMovieDetailResponse(omdbResponse),nil
}

func getGetMovieDetailResponse(spec movieDetailOMDB) *movie.GetMovieDetailResponse {
	resultChan:= make(chan *movie.Rating)

	for _,v := range spec.Ratings {
		go getRating(v, resultChan)
	}
	
	ratings := make([]*movie.Rating, len(spec.Ratings))
	for i := range ratings {
		ratings[i] = <-resultChan
	}

	getMovieDetailResponse := movie.GetMovieDetailResponse {
		Title : spec.Title,
		Year : spec.Year,
		Rated : spec.Rated,
		Released : spec.Released,
		Runtime : spec.Runtime,
		Genre : spec.Genre,
		Director : spec.Director,
		Writer : spec.Writer,
		Actors : spec.Actors,
		Plot : spec.Plot,
		Language : spec.Language,
		Country : spec.Country,
		Awards : spec.Awards,
		Poster : spec.Poster,
		Ratings : ratings,
		Metascore : spec.Metascore,
		ImdbRating : spec.ImdbRating,
		ImdbVotes : spec.ImdbVotes,
		ImdbID : spec.ImdbID,
		Type : spec.Type,
		DVD : spec.DVD,
		BoxOffice : spec.BoxOffice,
		Production : spec.Production,
		Website : spec.Website,
	}

	return &getMovieDetailResponse
}

func getRating(spec ratingOMDB, result chan *movie.Rating) {
	result <- &movie.Rating {
		Source : spec.Source,
		Value : spec.Value,
	}
}

func constructLoggerEntityFromSearchMovieRequest(spec *movie.SearchMovieRequest) *entity.LogEntity {
	return &entity.LogEntity{
		Timestamp : time.Now().UnixNano(),
		Request : spec.String(),
	}
}

func constructLoggerEntityFromGetMovieDetailRequest(spec *movie.GetMovieDetailRequest) *entity.LogEntity {
	return &entity.LogEntity{
		Timestamp : time.Now().UnixNano(),
		Request : spec.String(),
	}
}
