package impl

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/yecklilien/OMDB/logger/accessor"
	"github.com/yecklilien/OMDB/movie"
	"gorm.io/gorm"
)

//MovieAPI struct represent MovieAPI
type MovieAPI struct {
	httpClient http.Client
	logAcessor *accessor.LogAccessor
	omdbAPIKey string
}

//SearchMovie method
func (m *MovieAPI) SearchMovie(spec *movie.SearchMovieRequest) (*movie.SearchMovieResponse, error) {

	go m.logAcessor.Create(constructLoggerEntityFromSearchMovieRequest(spec))

	URL := constructSearchMovieGetURL(spec, m.omdbAPIKey)
	resp, err := m.httpClient.Get(URL)

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

//GetMovieDetail method
func (m *MovieAPI) GetMovieDetail(spec *movie.GetMovieDetailRequest) (*movie.GetMovieDetailResponse, error) {

	go m.logAcessor.Create(constructLoggerEntityFromGetMovieDetailRequest(spec))

	URL := constructGetMovieDetailURL(spec, m.omdbAPIKey)
	resp, err := m.httpClient.Get(URL)

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

func constructGetMovieDetailURL(spec *movie.GetMovieDetailRequest, apiKey string) string {
	object := map[string]string{
		"apikey": apiKey,
		"i":      spec.ImdbID,
	}
	return constructGetURL(object)
}

func constructSearchMovieGetURL(spec *movie.SearchMovieRequest, apiKey string) string {
	//https://medium.com/@felipedutratine/pass-environment-variables-from-docker-to-my-golang-2a967c5905fe
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
		URL += (key + "=" + value)
		i++
	}
	return URL
}

//NewMovieAPI function return new instance of MovieAPI
func NewMovieAPI(db *gorm.DB, omdbAPIKey string) *MovieAPI {

	timeout := time.Duration(5 * time.Second)
	httpClient := http.Client{
		Timeout: timeout,
	}

	logAccessor := accessor.NewLogAccessor(db)

	return &MovieAPI{
		httpClient: httpClient,
		logAcessor: logAccessor,
		omdbAPIKey: omdbAPIKey,
	}
}
