package impl

import (
	"context"
	"github.com/yecklilien/OMDB/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//MovieAPIGRPCServer to be implemented when register grpc
type MovieAPIGRPCServer struct{
	movie.UnimplementedMovieAPIServer
	movieAPI *MovieAPI
}

//NewMovieAPIGRPCServer function return new instance of MovieAPIGRPCServer
func NewMovieAPIGRPCServer(movieAPI *MovieAPI) *MovieAPIGRPCServer {
	return &MovieAPIGRPCServer{
		movieAPI : movieAPI,
	}
}

//SearchMovie grpc implementation
func (server *MovieAPIGRPCServer) SearchMovie(ctx context.Context, request *movie.SearchMovieRequest) (*movie.SearchMovieResponse, error) {
	response, err := server.movieAPI.SearchMovie(request)
	if err!=nil {
		return nil, status.Errorf(codes.Internal, "method SearchMovie got error" + err.Error())
	}
	return response,nil
} 

//GetMovieDetail grpc implementation
func (server *MovieAPIGRPCServer) GetMovieDetail(ctx context.Context, request *movie.GetMovieDetailRequest) (*movie.GetMovieDetailResponse, error) {
	response, err := server.movieAPI.GetMovieDetail(request)
	if err!=nil {
		return nil, status.Errorf(codes.Internal, "method GetMovieDetail got error" + err.Error())
	}
	return response,nil
} 
