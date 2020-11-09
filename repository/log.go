package repository

import (
	"log"

	"gorm.io/gorm"
	"github.com/yecklilien/OMDB/entity"
)

//Log struct containing LogEntity db method
type Log struct {
	db *gorm.DB
}

//NewLog initialize Log
func NewLog(db *gorm.DB) *Log {
	db.AutoMigrate(&entity.Log{})
	return &Log{
		db: db,
	}
}

//Create LogEntity into db
func (accessor *Log) Create(entity entity.Log) error {
	result:= accessor.db.Create(&entity)
	if result.Error != nil {
		log.Printf("error when create Log : %v", result.Error);
		return result.Error
	}
	return nil;
} 

//Get LogEntity from db by id
func (accessor *Log) Get(id int64) (*entity.Log, error) {
	var logEntity entity.Log
	result:= accessor.db.First(&logEntity,id)
	if result.Error != nil {
		log.Printf("error when get LogEntity : %v", result.Error);
		return nil,result.Error
	}
	return &logEntity,nil
}