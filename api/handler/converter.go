package handler

import(
	"log"
	"strconv"

	"github.com/yecklilien/OMDB/api/presenter"
	"github.com/yecklilien/OMDB/entity"
)

func constructSearchMovieResponse(spec entity.SearchMovieResult) (*presenter.SearchMovieResponse, error) {
	
	resultChan:= make(chan *presenter.Movie)

	for _,v := range spec.Movies {
		go constructMovie(v, resultChan)
	}
	
	movies := make([]*presenter.Movie, len(spec.Movies))
	for i := range movies {
		movies[i] = <-resultChan
	}

	totalResults,err := strconv.Atoi(spec.TotalResult);

	if(err!=nil) {
		log.Print(err)
		return nil,err;
	}

	searchMovieResponse := presenter.SearchMovieResponse {
		Movies: movies,
		TotalResult: int32(totalResults),
	}

	return &searchMovieResponse,nil
}

func constructMovie(spec entity.Movie, result chan *presenter.Movie){
	result <- &presenter.Movie {
		Title : spec.Title,
		Year : spec.Year,
		ImdbID : spec.ImdbID,
		Poster : spec.Poster,
		Type : spec.Type,
	}
}

func constructGetMovieDetailResponse(spec entity.MovieDetail) *presenter.GetMovieDetailResponse {
	resultChan:= make(chan *presenter.Rating)

	for _,v := range spec.Ratings {
		go constructRating(v, resultChan)
	}
	
	ratings := make([]*presenter.Rating, len(spec.Ratings))
	for i := range ratings {
		ratings[i] = <-resultChan
	}

	getMovieDetailResponse := presenter.GetMovieDetailResponse {
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

func constructRating(spec entity.Rating, result chan *presenter.Rating) {
	result <- &presenter.Rating {
		Source : spec.Source,
		Value : spec.Value,
	}
}