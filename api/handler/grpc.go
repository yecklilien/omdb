package handler

import (
	"context"
	"time"

	"github.com/yecklilien/OMDB/usecase/movie"
	"github.com/yecklilien/OMDB/usecase/log"
	"github.com/yecklilien/OMDB/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/yecklilien/OMDB/api/presenter"
)

//GRPCHandler to be implemented when register grpc
type GRPCHandler struct{
	presenter.UnimplementedMovieAPIServer
	movieUsecase movie.UseCase
	logUsecase log.UseCase
}

//NewGRPCHandler function return new instance of GRPCHandler
func NewGRPCHandler(movieUsecase movie.UseCase, logUsecase log.UseCase) *GRPCHandler {
	return &GRPCHandler{
		movieUsecase : movieUsecase,
		logUsecase : logUsecase,
	}
}

//SearchMovie grpc implementation
func (handler *GRPCHandler) SearchMovie(ctx context.Context, request *presenter.SearchMovieRequest) (*presenter.SearchMovieResponse, error) {
	
	go handler.logUsecase.Log(entity.Log {
		Timestamp : time.Now().UnixNano(),
		Request : request.String(),
	});

	searchMovieResult, err := handler.movieUsecase.SearchMovie(
		entity.SearchMovieSpec{
			Page : request.Page,
			Query : request.Query,
		})

	if err!=nil {
		return nil, status.Errorf(codes.Internal, "method SearchMovie got error" + err.Error())
	}
	
	response, err:=constructSearchMovieResponse(*searchMovieResult)

	if err!=nil {
		return nil, status.Errorf(codes.Internal, "method SearchMovie got error" + err.Error())
	}
	
	return response, err
} 

//GetMovieDetail grpc implementation
func (handler *GRPCHandler) GetMovieDetail(ctx context.Context, request *presenter.GetMovieDetailRequest) (*presenter.GetMovieDetailResponse, error) {
	go handler.logUsecase.Log(entity.Log {
		Timestamp : time.Now().UnixNano(),
		Request : request.String(),
	});
	
	movieDetail, err := handler.movieUsecase.GetMovieDetail(entity.GetMovieDetailSpec {
		ImdbID : request.ImdbID,
	});
	
	if err!=nil {
		return nil, status.Errorf(codes.Internal, "method GetMovieDetail got error" + err.Error())
	}
	
	return constructGetMovieDetailResponse(*movieDetail), err
} 
