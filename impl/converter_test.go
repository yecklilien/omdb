package impl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
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

func constructExpectedSearchMovieResponse() *movie.SearchMovieResponse {
	file, _ := ioutil.ReadFile("test_file/SearchMovieResponse.json")
	result := &movie.SearchMovieResponse{}
	_ = json.Unmarshal([]byte(file), result)
	return result
}

func constructSearchMovieResponseOMDB() *searchMovieResponseOMDB {
	file, _ := ioutil.ReadFile("test_file/SearchMovieResponseOMDB.json")
	result := &searchMovieResponseOMDB{}
	_ = json.Unmarshal([]byte(file), result)
	return result
}

func TestConstructSearchMovieResponse(t *testing.T) {
	file, _ := ioutil.ReadFile("test_file/SearchMovieResponseOMDB.json")
	searchMovieResponse,err := constructSearchMovieResponse([]byte(file))
	if err!= nil {
		t.Fatal("Unexpected error thrown")
	}

	sort.SliceStable(searchMovieResponse.Movies, func(i, j int) bool {
		return searchMovieResponse.Movies[i].Title < searchMovieResponse.Movies[j].Title
	})

	expectedSearchMovieResponse := constructExpectedSearchMovieResponse()

	sort.SliceStable(expectedSearchMovieResponse.Movies, func(i, j int) bool {
		return expectedSearchMovieResponse.Movies[i].Title < expectedSearchMovieResponse.Movies[j].Title
	})

	assertEqual(t,searchMovieResponse.String(),expectedSearchMovieResponse.String(),"")
}

func TestGetSearchMovieResponse(t *testing.T) {
	searchMovieResponseOMDB := constructSearchMovieResponseOMDB()
	expectedSearchMovieResponse := constructExpectedSearchMovieResponse()
	
	sort.SliceStable(expectedSearchMovieResponse.Movies, func(i, j int) bool {
		return expectedSearchMovieResponse.Movies[i].Title < expectedSearchMovieResponse.Movies[j].Title
	})

	actualSearchMovieResponse,err := getSearchMovieResponse(*searchMovieResponseOMDB)

	sort.SliceStable(actualSearchMovieResponse.Movies, func(i, j int) bool {
		return actualSearchMovieResponse.Movies[i].Title < actualSearchMovieResponse.Movies[j].Title
	})
	
	if err != nil {
		t.Fatal("Unexpected error when get search movie response")
	}

	assertEqual(t, expectedSearchMovieResponse.String(), actualSearchMovieResponse.String(), "" )
}

func TestGetMovie(t *testing.T) {
	movieOMDB := constructSearchMovieResponseOMDB().Movies[0]
	resultChan := make(chan *movie.Movie)
	go getMovie(movieOMDB, resultChan)
	var movie *movie.Movie
	movie = <-resultChan

	assertEqual(t, movieOMDB.ImdbID, movie.ImdbID, "")
	assertEqual(t, movieOMDB.Poster, movie.Poster, "")
	assertEqual(t, movieOMDB.Title, movie.Title, "")
	assertEqual(t, movieOMDB.Type, movie.Type, "")
	assertEqual(t, movieOMDB.Year, movie.Year, "")
}

func TestConstructGetMovieDetailResponse(t *testing.T) {
	expectedMovieDetail := constructExpectedGetMovieDetailResponse()
	sort.SliceStable(expectedMovieDetail.Ratings, func(i, j int) bool {
		if expectedMovieDetail.Ratings[i].Source == expectedMovieDetail.Ratings[j].Source {
			return expectedMovieDetail.Ratings[i].Value < expectedMovieDetail.Ratings[j].Value
		}
		return expectedMovieDetail.Ratings[i].Source < expectedMovieDetail.Ratings[j].Source
	})

	file, err := ioutil.ReadFile("test_file/MovieDetailOMDB.json")

	if err != nil {
		fmt.Println(err)
		t.Fatalf("error when construct Get Movie Detail Response: %v", err)
	}

	actualMovieDetail, err := constructGetMovieDetailResponse([]byte(file))

	fmt.Println(actualMovieDetail)

	if err != nil {
		t.Fatalf("error when construct Get Movie Detail Response: %v", err)
	}

	sort.SliceStable(actualMovieDetail.Ratings, func(i, j int) bool {
		if actualMovieDetail.Ratings[i].Source == actualMovieDetail.Ratings[j].Source {
			return actualMovieDetail.Ratings[i].Value < actualMovieDetail.Ratings[j].Value
		}
		return actualMovieDetail.Ratings[i].Source < actualMovieDetail.Ratings[j].Source
	})

	assertEqual(t, expectedMovieDetail.String(), actualMovieDetail.String(), "")
}

func TestGetGetMovieDetailResponse(t *testing.T) {
	movieDetailOMDB := constructMovieDetailOMDB()
	expectedMovieDetail := constructExpectedGetMovieDetailResponse()

	sort.SliceStable(expectedMovieDetail.Ratings, func(i, j int) bool {
		if expectedMovieDetail.Ratings[i].Source == expectedMovieDetail.Ratings[j].Source {
			return expectedMovieDetail.Ratings[i].Value < expectedMovieDetail.Ratings[j].Value
		}
		return expectedMovieDetail.Ratings[i].Source < expectedMovieDetail.Ratings[j].Source
	})

	movieDetail := getGetMovieDetailResponse(*movieDetailOMDB)

	sort.SliceStable(movieDetail.Ratings, func(i, j int) bool {
		if movieDetail.Ratings[i].Source == movieDetail.Ratings[j].Source {
			return movieDetail.Ratings[i].Value < movieDetail.Ratings[j].Value
		}
		return movieDetail.Ratings[i].Source < movieDetail.Ratings[j].Source
	})

	assertEqual(t, expectedMovieDetail.String(), movieDetail.String(), "")
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
