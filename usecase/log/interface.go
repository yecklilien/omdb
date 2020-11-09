package log

import (
	"github.com/yecklilien/OMDB/entity"
)

type Repository interface {
	Create(e entity.Log) error
	Get(id int64) (*entity.Log,error)
}

type UseCase interface {
	Log(e entity.Log) error
}
