package main

import (
	"context"
	"log"

	"github.com/yecklilien/OMDB/movie"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := movie.NewMovieAPIClient(conn)

	searchMovieResponse, err := client.SearchMovie(context.Background(), constructSearchMovieRequest())
	if err != nil {
		log.Fatalf("SearchMovie failed: %v", err)
	}
	log.Printf("SearchMovieResponse: %v", searchMovieResponse)

	getMovieDetailResponse, err := client.GetMovieDetail(context.Background(), constructGetMovieDetailRequest())
	if err != nil {
		log.Fatalf("GetMovieDetail failed: %v", err)
	}
	log.Printf("GetMovieDetailResponse: %v", getMovieDetailResponse)
}

func constructSearchMovieRequest() *movie.SearchMovieRequest {
	return &movie.SearchMovieRequest{
		Page:  1,
		Query: "Batman",
	}
}

func constructGetMovieDetailRequest() *movie.GetMovieDetailRequest {
	return &movie.GetMovieDetailRequest{
		ImdbID: "tt2166834",
	}
}
