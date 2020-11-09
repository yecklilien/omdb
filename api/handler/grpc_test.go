package handler

import (
	"errors"
	"testing"
	"context"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"sort"

	"github.com/yecklilien/OMDB/entity"
	"github.com/yecklilien/OMDB/api/presenter"
)

type mockMovieUsecase struct {}

func (m *mockMovieUsecase) GetMovieDetail(spec entity.GetMovieDetailSpec) (*entity.MovieDetail,error) {
	return mockGetMovieDetail(spec)
}

func (m *mockMovieUsecase) SearchMovie (spec entity.SearchMovieSpec) (*entity.SearchMovieResult,error) {
	return mockSearchMovie(spec)
}

type mockLogUsecase struct {}

func (m *mockLogUsecase) Log (e entity.Log) error {
	return nil
}

var mockGetMovieDetail func (entity.GetMovieDetailSpec) (*entity.MovieDetail,error)
var mockSearchMovie func (entity.SearchMovieSpec) (*entity.SearchMovieResult,error)

var grpcHandler GRPCHandler = *NewGRPCHandler(&mockMovieUsecase{},&mockLogUsecase{})

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestGRPCSearchMovie(t *testing.T) {
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

	searchMovieResponse, err := grpcHandler.SearchMovie(
		context.Background(),
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		});
	
	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(searchMovieResponse.Movies, func(i, j int) bool {
		return searchMovieResponse.Movies[i].Title < searchMovieResponse.Movies[j].Title
	})

	assertEqual(t,expectedResult.String(),searchMovieResponse.String(),"")
}

func TestGRPCSearchMovie_errorSearchMovie(t *testing.T) {
	mockSearchMovie = func (e entity.SearchMovieSpec) (*entity.SearchMovieResult,error) {
		return nil,errors.New("Expected Error Thrown");
	}

	_, err := grpcHandler.SearchMovie(
		context.Background(),
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		});
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}

func TestGRPCSearchMovie_errorConvert(t *testing.T) {
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

	_, err = grpcHandler.SearchMovie(
		context.Background(),
		&presenter.SearchMovieRequest{
			Page : 1,
			Query : "The Prestige",
		});
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}

func TestGRPCGetMovieDetail(t *testing.T) {
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

	getMovieDetailResponse , err := grpcHandler.GetMovieDetail(
		context.Background(),
		&presenter.GetMovieDetailRequest{
			ImdbID : "tt0482571",
		});
	
	if err != nil {
		t.Fatalf("Unexpected error thrown: %v",err)
	}

	sort.SliceStable(getMovieDetailResponse.Ratings, func(i, j int) bool {
		return getMovieDetailResponse.Ratings[i].Source < getMovieDetailResponse.Ratings[j].Source
	})

	assertEqual(t,expectedResult.String(),getMovieDetailResponse.String(),"")
}

func TestGRPCGetMovieDetail_errorGetMovieDetail(t *testing.T) {
	
	mockGetMovieDetail = func (e entity.GetMovieDetailSpec) (*entity.MovieDetail,error) {
		return nil,errors.New("Expected error thrown");
	}

	_ , err := grpcHandler.GetMovieDetail(
		context.Background(),
		&presenter.GetMovieDetailRequest{
			ImdbID : "tt0482571",
		});
	
	if err == nil {
		t.Fatalf("Unexpected no error thrown")
	}
}