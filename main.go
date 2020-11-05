package main

import (
	"fmt"
	"github.com/yecklilien/OMDB/model"
	"github.com/yecklilien/OMDB/impl"
)

func main() {
	movieAPI:= impl.NewMovieAPI()
	searchMovieRequest := model.SearchMovieRequest {
		Page : 1,
		Query : "Batman",
	}
	searchMovieResponse,err := movieAPI.SearchMovie(searchMovieRequest);
	if(err != nil) {
		fmt.Println(err)
	}
	fmt.Println(searchMovieResponse.Movies[0])

	getMovieDetailRequest := model.GetMovieDetailRequest {
		ImdbID : "tt2313197",
	}

	getMovieDetailResponse,err := movieAPI.GetMovieDetail(getMovieDetailRequest);
	if(err != nil) {
		fmt.Println(err)
	}
	fmt.Println(getMovieDetailResponse);
}