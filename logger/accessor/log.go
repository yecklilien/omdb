package accessor

import (
	"log"

	"gorm.io/gorm"
	"github.com/yecklilien/OMDB/logger/entity"
)

//LogAccessor struct containing LogEntity db method
type LogAccessor struct {
	db *gorm.DB
}

//NewLogAccessor initialize LogAccessor
func NewLogAccessor(db *gorm.DB) *LogAccessor {
	db.AutoMigrate(&entity.LogEntity{})
	return &LogAccessor{
		db: db,
	}
}

//Create LogEntity into db
func (accessor *LogAccessor) Create(entity *entity.LogEntity) error {
	result:= accessor.db.Create(&entity)
	if result.Error != nil {
		log.Fatalf("error when create LogEntity : %v", result.Error);
		return result.Error
	}
	return nil;
} 

//Get LogEntity from db by id
func (accessor *LogAccessor) Get(id *int64) (*entity.LogEntity, error) {
	var logEntity entity.LogEntity
	result:= accessor.db.First(&logEntity,id)
	if result.Error != nil {
		log.Fatalf("error when get LogEntity : %v", result.Error);
		return nil,result.Error
	}
	return &logEntity,nil
}
