package repository

import (
	"log"

	"gorm.io/gorm"
	"github.com/yecklilien/OMDB/entity"
)

//LogAccessor struct containing LogEntity db method
type LogAccessor struct {
	db *gorm.DB
}

//NewLogAccessor initialize LogAccessor
func NewLogAccessor(db *gorm.DB) *LogAccessor {
	db.AutoMigrate(&entity.Log{})
	return &LogAccessor{
		db: db,
	}
}

//Create LogEntity into db
func (accessor *LogAccessor) Create(entity *entity.Log) error {
	result:= accessor.db.Create(&entity)
	if result.Error != nil {
		log.Fatalf("error when create Log : %v", result.Error);
		return result.Error
	}
	return nil;
} 

//Get LogEntity from db by id
func (accessor *LogAccessor) Get(id *int64) (*entity.Log, error) {
	var logEntity entity.Log
	result:= accessor.db.First(&logEntity,id)
	if result.Error != nil {
		log.Fatalf("error when get LogEntity : %v", result.Error);
		return nil,result.Error
	}
	return &logEntity,nil
}