package movie

import (
	"github.com/yecklilien/OMDB/entity"
)

//UseCase movie
type UseCase interface {
	SearchMovie(spec entity.SearchMovieSpec) (* entity.SearchMovieResult, error)
	GetMovieDetail(spec entity.GetMovieDetailSpec) (* entity.MovieDetail, error)
}
