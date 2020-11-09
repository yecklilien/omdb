package log

import (
	"errors"
	"testing"
	
	"github.com/yecklilien/OMDB/entity"
)

type Repo struct {
}

func (r *Repo) Create(e entity.Log) error {
	return createMock(e);
}

func (r *Repo) Get(int64) *entity.Log{
	return nil;
}

var createMock func(e entity.Log) error
var service Service

func init() {
	service = *NewService(&Repo{})
}

func TestLog(t *testing.T) {
	createMock = func(e entity.Log) error {
		return nil;
	}
	
	err := service.Log(entity.Log{})
	if err != nil {
		t.Fatalf("Unexpected Error")
	}
}

func TestLogError(t *testing.T) {
	createMock = func(e entity.Log) error {
		return errors.New("Error");
	}
	err := service.Log(entity.Log{})
	if err == nil {
		t.Fatalf("Expected Exception")
	}
}



