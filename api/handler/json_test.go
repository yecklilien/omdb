package handler

import (
	"net/http"
	"errors"
	"testing"
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/yecklilien/OMDB/entity"
	"github.com/yecklilien/OMDB/api/presenter"
)

var jsonHandler JSONHTTPHandler = *NewJSONHTTPHandler(&mockMovieUsecase{},&mockLogUsecase{})

func TestJSONSearchMovie(t *testing.T) {
	file, _ := ioutil.ReadFile("test_file/PresenterSearchMovieResponse.json")
	expectedResult := &presenter.SearchMovieResponse{}
	err:=json.Unmarshal([]byte(file), expectedResult)

	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(expectedResult.Movies, func(i, j int) bool {
		return expectedResult.Movies[i].Title < expectedResult.Movies[j].Title
	})

	mockSearchMovie = func (e entity.SearchMovieSpec) (*entity.SearchMovieResult,error) {
		file, _ := ioutil.ReadFile("test_file/SearchMovieResult.json")
		result := &entity.SearchMovieResult{}
		err:=json.Unmarshal([]byte(file), result)
		if e.Page == 1 && e.Query == "The Prestige" {
			return result,err;
		}
		return nil,err;
	}
	var searchMovieResponse presenter.SearchMovieResponse
	err = jsonHandler.SearchMovie(
		&http.Request{},
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		},
		&searchMovieResponse);
	
	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(searchMovieResponse.Movies, func(i, j int) bool {
		return searchMovieResponse.Movies[i].Title < searchMovieResponse.Movies[j].Title
	})

	assertEqual(t,expectedResult.String(),searchMovieResponse.String(),"")
}

func TestSearchMovie_errorSearchMovie(t *testing.T) {
	mockSearchMovie = func (e entity.SearchMovieSpec) (*entity.SearchMovieResult,error) {
		return nil,errors.New("Expected Error Thrown");
	}

	var searchMovieResponse presenter.SearchMovieResponse
	err := jsonHandler.SearchMovie(
		&http.Request{},
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		},
		&searchMovieResponse);
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}

func TestSearchMovie_errorConvert(t *testing.T) {
	file, _ := ioutil.ReadFile("test_file/PresenterSearchMovieResponse.json")
	expectedResult := &presenter.SearchMovieResponse{}
	err:=json.Unmarshal([]byte(file), expectedResult)

	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(expectedResult.Movies, func(i, j int) bool {
		return expectedResult.Movies[i].Title < expectedResult.Movies[j].Title
	})

	mockSearchMovie = func (e entity.SearchMovieSpec) (*entity.SearchMovieResult,error) {
		file, _ := ioutil.ReadFile("test_file/SearchMovieResult.json")
		result := &entity.SearchMovieResult{}
		err:=json.Unmarshal([]byte(file), result)
		if e.Page == 1 && e.Query == "The Prestige" {
			result.TotalResult = "Error"
			return result,err;
		}
		return nil,err;
	}

	var searchMovieResponse presenter.SearchMovieResponse
	err = jsonHandler.SearchMovie(
		&http.Request{},
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		},
		&searchMovieResponse);
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}

func TestGetMovieDetail(t *testing.T) {
	file, _ := ioutil.ReadFile("test_file/MovieDetail.json")
	expectedResult := &presenter.GetMovieDetailResponse{}
	err:=json.Unmarshal([]byte(file), expectedResult)

	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(expectedResult.Ratings, func(i, j int) bool {
		return expectedResult.Ratings[i].Source < expectedResult.Ratings[j].Source
	})

	mockGetMovieDetail = func (e entity.GetMovieDetailSpec) (*entity.MovieDetail,error) {
		file, _ := ioutil.ReadFile("test_file/MovieDetail.json")
		result := &entity.MovieDetail{}
		err:=json.Unmarshal([]byte(file), result)
		if e.ImdbID == "tt0482571" {
			return result,err;
		}
		return nil,err;
	}

	var getMovieDetailResponse presenter.GetMovieDetailResponse

	err = jsonHandler.GetMovieDetail(
		&http.Request{},
		&presenter.GetMovieDetailRequest{
			ImdbID : "tt0482571",
		},
		&getMovieDetailResponse);
	
	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(getMovieDetailResponse.Ratings, func(i, j int) bool {
		return getMovieDetailResponse.Ratings[i].Source < getMovieDetailResponse.Ratings[j].Source
	})

	assertEqual(t,expectedResult.String(),getMovieDetailResponse.String(),"")
}

func TestGetMovieDetail_errorGetMovieDetail(t *testing.T) {
	
	mockGetMovieDetail = func (e entity.GetMovieDetailSpec) (*entity.MovieDetail,error) {
		return nil,errors.New("Expected error thrown");
	}

	var getMovieDetailResponse presenter.GetMovieDetailResponse

	err := jsonHandler.GetMovieDetail(
		&http.Request{},
		&presenter.GetMovieDetailRequest{
			ImdbID : "tt0482571",
		},
		&getMovieDetailResponse);
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}