package handler

import (
	"net/http"
	"time"
	logger "log"

	"github.com/yecklilien/OMDB/api/presenter"
	"github.com/yecklilien/OMDB/usecase/movie"
	"github.com/yecklilien/OMDB/usecase/log"
	"github.com/yecklilien/OMDB/entity"
)

//JSONHTTPHandler to be implemented
type JSONHTTPHandler struct {
	movieUsecase movie.UseCase
	logUsecase log.UseCase
}

//NewJSONHTTPHandler function return new instance of JSONHTTPHandler
func NewJSONHTTPHandler(movieUsecase movie.UseCase, logUsecase log.UseCase) *JSONHTTPHandler {
	return &JSONHTTPHandler{
		movieUsecase: movieUsecase,
		logUsecase: logUsecase,
	}
}

//SearchMovie http json rpc implementation
func (handler *JSONHTTPHandler) SearchMovie(r *http.Request, spec *presenter.SearchMovieRequest, result *presenter.SearchMovieResponse) error {
	
	go handler.logUsecase.Log(entity.Log {
		Timestamp : time.Now().UnixNano(),
		Request : spec.String(),
	});

	searchMovieResult, err := handler.movieUsecase.SearchMovie(
		entity.SearchMovieSpec{
			Page : spec.Page,
			Query : spec.Query,
		})

	if err!=nil {
		logger.Printf("method SearchMovie got error : %v", err)
		return err
	}
	
	response, err:=constructSearchMovieResponse(*searchMovieResult)

	if err!=nil {
		logger.Printf("method SearchMovie got error : %v", err)
		return err
	}

	*result = *response
	return err
}

//GetMovieDetail http json rpc implementation
func (handler *JSONHTTPHandler) GetMovieDetail(r *http.Request, spec *presenter.GetMovieDetailRequest, result *presenter.GetMovieDetailResponse) error {
	go handler.logUsecase.Log(entity.Log {
		Timestamp : time.Now().UnixNano(),
		Request : spec.String(),
	});
	
	movieDetail, err := handler.movieUsecase.GetMovieDetail(entity.GetMovieDetailSpec {
		ImdbID : spec.ImdbID,
	});
	
	if err!=nil {
		logger.Printf("method GetMovieDetail got error : %v", err)
		return err
	}
	
	*result = *constructGetMovieDetailResponse(*movieDetail)
	return err
}
