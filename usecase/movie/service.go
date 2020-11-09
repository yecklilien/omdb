package movie

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
	"log"
	"errors"
	"io/ioutil"
	"encoding/json"


	"github.com/yecklilien/OMDB/entity"
)

//Service movie
type Service struct {
	httpClient httpClient
	omdbAPIKey string
}

type httpClient interface {
	Get(string) (resp *http.Response, err error) 
}

//SearchMovie usecase
func (service *Service) SearchMovie(spec entity.SearchMovieSpec) (* entity.SearchMovieResult, error) {

	URL := constructSearchMovieGetURL(spec, service.omdbAPIKey)
	resp, err := service.httpClient.Get(URL)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	response, err := constructSearchMovieResponse(body)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return response, nil
}

//GetMovieDetail usecase
func (service *Service) GetMovieDetail(spec entity.GetMovieDetailSpec) (* entity.MovieDetail, error) {
	URL := constructGetMovieDetailURL(spec, service.omdbAPIKey)
	resp, err := service.httpClient.Get(URL)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	response, err := constructGetMovieDetailResponse(body)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return response, nil
}

func constructGetMovieDetailURL(spec entity.GetMovieDetailSpec, apiKey string) string {
	object := map[string]string{
		"apikey": apiKey,
		"i":      spec.ImdbID,
	}
	return constructGetURL(object)
}

func constructSearchMovieGetURL(spec entity.SearchMovieSpec, apiKey string) string {
	object := map[string]string{
		"apikey": apiKey,
		"s":      spec.Query,
		"page":   strconv.Itoa(int(spec.Page)),
	}
	return constructGetURL(object)
}

func constructGetURL(object map[string]string) string {
	URL := "http://www.omdbapi.com/"
	i := 1
	for key, value := range object {
		if i == 1 {
			URL += "?"
		} else {
			URL += "&"
		}
		URL += (key + "=" + url.QueryEscape(value))
		i++
	}
	return URL
}

//NewService return new movie service
func NewService(omdbAPIKey string) *Service{

	timeout := time.Duration(5 * time.Second)
	httpClient := http.Client{
		Timeout: timeout,
	}

	return newService(&httpClient,omdbAPIKey)
}

func newService(httpClient httpClient, omdbAPIKey string) *Service {
	return &Service{
		httpClient: httpClient,
		omdbAPIKey: omdbAPIKey,
	}
}

func constructSearchMovieResponse(bytes []byte) (*entity.SearchMovieResult, error) {
	var searchMovieResult entity.SearchMovieResult
	err:= json.Unmarshal(bytes,&searchMovieResult)
	
	if err!=nil {
		log.Print(err);
		return nil,err
	}

	if searchMovieResult.Response != "True" {
		log.Print(searchMovieResult.Error);
		return nil, errors.New(searchMovieResult.Error)
	}
	
	return &searchMovieResult,nil
}

func constructGetMovieDetailResponse(bytes []byte) (* entity.MovieDetail, error) {
	var  movieDetail entity.MovieDetail
	err:= json.Unmarshal(bytes,&movieDetail)
	
	if err!=nil {
		log.Print(err);
		return nil,err
	}

	if movieDetail.Response != "True" {
		log.Print(movieDetail.Error);
		return nil, errors.New(movieDetail.Error)
	}

	return &movieDetail,nil
}

