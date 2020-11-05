package impl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/yecklilien/OMDB/movie"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func constructMovieDetailOMDB() *movieDetailOMDB {
	file, _ := ioutil.ReadFile("test_file/MovieDetailOMDB.json")
	result := &movieDetailOMDB{}
	_ = json.Unmarshal([]byte(file), result)
	return result
}

func constructExpectedGetMovieDetailResponse() *movie.GetMovieDetailResponse {
	file, _ := ioutil.ReadFile("test_file/GetMovieDetailResponse.json")
	result := &movie.GetMovieDetailResponse{}
	_ = json.Unmarshal([]byte(file), result)
	return result
}

func TestGetGetMovieDetailResponse(t *testing.T) {
	// movieDetailOMDB := constructMovieDetailOMDB()
	// expectedMovieDetail := constructExpectedGetMovieDetailResponse()
	// movieDetail := getGetMovieDetailResponse(*movieDetailOMDB)
	// assertEqual(t, expectedMovieDetail, movieDetail, "")
}

func TestGetRating(t *testing.T) {
	omdbRating := constructMovieDetailOMDB().Ratings[0]
	resultChan := make(chan *movie.Rating)
	go getRating(omdbRating, resultChan)
	var rating *movie.Rating
	rating = <-resultChan

	assertEqual(t, omdbRating.Source, rating.Source, "")
	assertEqual(t, omdbRating.Value, rating.Value, "")
}

func TestConstructLoggerEntityFromSearchMovieRequest(t *testing.T) {
	spec := movie.SearchMovieRequest{
		Page:  1,
		Query: "Batman",
	}

	logEntity := constructLoggerEntityFromSearchMovieRequest(&spec)

	if logEntity.Timestamp == 0 {
		t.Fatal("LogEntity Timestamp is 0")
	}

	assertEqual(t, logEntity.Request, spec.String(), "")
}

func TestConstructLoggerEntityFromGetMovieDetailRequest(t *testing.T) {
	spec := movie.GetMovieDetailRequest{
		ImdbID: "testId",
	}

	logEntity := constructLoggerEntityFromGetMovieDetailRequest(&spec)

	if logEntity.Timestamp == 0 {
		t.Fatal("LogEntity Timestamp is 0")
	}

	assertEqual(t, logEntity.Request, spec.String(), "")
}
