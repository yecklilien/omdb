package movie

import (
	"fmt"
	"bytes"
	"testing"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"

	"github.com/yecklilien/OMDB/entity"
)

type mockHttpClient struct {}

func (m*mockHttpClient) Get(urlString string) (resp *http.Response, err error) {
	return mockHttpClientGet(urlString)
}

type errReadCloser struct{}

func (errReadCloser) Read(p []byte) (n int, err error) {
    return 0, errors.New("test error")
}
func (errReadCloser) Close() (err error) {
    return errors.New("test error")
}

var mockHttpClientGet func(string) (*http.Response,error)
 
var service Service = *newService(&mockHttpClient{},"testKey")


func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestNewService(t *testing.T) {
	NewService("testKey");
}

func TestSearchMovie(t *testing.T){
	file, _ := ioutil.ReadFile("test_file/SearchMovieResponse.json")
	expectedResult := &entity.SearchMovieResult{}
	err:= json.Unmarshal([]byte(file), expectedResult)

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		URL,_ := url.Parse(urlString)
		Query, _ := url.ParseQuery(URL.RawQuery)
		if Query.Get("apikey") == "testKey" && Query.Get("s") == "The Prestige" &&  Query.Get("page") == "1" {
			
			response := http.Response{
				Body : ioutil.NopCloser(bytes.NewBuffer([]byte(file))),
			}
			return &response,nil
		}
		return nil,errors.New("UrlString was not mocked");
	}

	searchMovieResult,err:= service.SearchMovie(entity.SearchMovieSpec{
		Page : 1,
		Query : "The Prestige",
	})

	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t,fmt.Sprint(searchMovieResult),fmt.Sprint(expectedResult),"")
}

func TestSearchMovie_errorGet(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		return nil,errors.New("Error get");
	}

	_,err:= service.SearchMovie(entity.SearchMovieSpec{
		Page : 1,
		Query : "The Prestige",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestSearchMovie_errorRead(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		response := http.Response{
			Body : errReadCloser{},
		}
		return &response,nil
	}

	_,err:= service.SearchMovie(entity.SearchMovieSpec{
		Page : 1,
		Query : "The Prestige",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestSearchMovie_errorUnmarshal(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		response := http.Response{
			Body : ioutil.NopCloser(bytes.NewBuffer([]byte{})),
		}
		return &response,nil
	}

	_,err:= service.SearchMovie(entity.SearchMovieSpec{
		Page : 1,
		Query : "The Prestige",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestSearchMovie_ErrorResult(t *testing.T){
	file, _ := ioutil.ReadFile("test_file/ErrorResponse.json")
	expectedResult := &entity.SearchMovieResult{}
	_ = json.Unmarshal([]byte(file), expectedResult)

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		URL,_ := url.Parse(urlString)
		Query, _ := url.ParseQuery(URL.RawQuery)
		if Query.Get("apikey") == "testKey" && Query.Get("s") == "The Prestige" &&  Query.Get("page") == "1" {	
			response := http.Response{
				Body : ioutil.NopCloser(bytes.NewBuffer([]byte(file))),
			}
			return &response,nil
		}
		return nil,errors.New("UrlString was not mocked");
	}

	_,err:= service.SearchMovie(entity.SearchMovieSpec{
		Page : 1,
		Query : "The Prestige",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	}
}

func TestGetMovieDetail(t *testing.T){
	file, _ := ioutil.ReadFile("test_file/MovieDetail.json")
	expectedResult := &entity.MovieDetail{}
	_ = json.Unmarshal([]byte(file), expectedResult)

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		URL,_ := url.Parse(urlString)
		Query, _ := url.ParseQuery(URL.RawQuery)
		if Query.Get("apikey") == "testKey" && Query.Get("i") == "tt0482571" {
			response := http.Response{
				Body : ioutil.NopCloser(bytes.NewBuffer([]byte(file))),
			}
			return &response,nil
		}
		return nil,fmt.Errorf("UrlString was not mocked %v",urlString);
	}

	movieDetail,err:= service.GetMovieDetail(entity.GetMovieDetailSpec{
		ImdbID : "tt0482571",
	})

	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t,fmt.Sprint(movieDetail),fmt.Sprint(expectedResult),"")
}

func TestGetMovieDetail_errorGet(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		return nil,errors.New("Error get");
	}

	_,err:= service.GetMovieDetail(entity.GetMovieDetailSpec{
		ImdbID : "tt0482571",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestGetMovieDetail_errorRead(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		response := http.Response{
			Body : errReadCloser{},
		}
		return &response,nil
	}

	_,err:= service.GetMovieDetail(entity.GetMovieDetailSpec{
		ImdbID : "tt0482571",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestGetMovieDetail_errorUnmarshal(t *testing.T){

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		response := http.Response{
			Body : ioutil.NopCloser(bytes.NewBuffer([]byte{})),
		}
		return &response,nil
	}

	_,err:= service.GetMovieDetail(entity.GetMovieDetailSpec{
		ImdbID : "tt0482571",
	})

	if err == nil {
		t.Fatal("Error was expected but none thrown")
	} 
}

func TestGetMovieDetail_ErrorResult(t *testing.T){
	file, _ := ioutil.ReadFile("test_file/ErrorResponse.json")
	expectedResult := &entity.SearchMovieResult{}
	_ = json.Unmarshal([]byte(file), expectedResult)

	mockHttpClientGet = func(urlString string) (*http.Response,error) {
		URL,_ := url.Parse(urlString)
		Query, _ := url.ParseQuery(URL.RawQuery)
		if Query.Get("apikey") == "testKey" && Query.Get("i") == "tt0482571" {
			response := http.Response{
				Body : ioutil.NopCloser(bytes.NewBuffer([]byte(file))),
			}
			return &response,nil
		}
		return nil,errors.New("UrlString was not mocked");
	}
	_,err:= service.GetMovieDetail(entity.GetMovieDetailSpec{
		ImdbID : "tt0482571",
	})
	
	if err == nil {
		t.Fatal("Error was expected but none thrown")
	}
}