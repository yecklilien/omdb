package log

import (
	"log"

	"github.com/yecklilien/OMDB/entity"
)

//Service log
type Service struct {
	repo Repository
}

//Log logging service
func (s *Service) Log(e entity.Log) error {
	err := s.repo.Create(e)
	if err != nil {
		log.Printf("error when logging with spec : %v, error: %v", e, err)
	}
	return err
}

//NewService book service
func NewService(repo Repository) *Service {
	return &Service{
		repo : repo,
	}
}