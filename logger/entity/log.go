package entity

//LogEntity entity
type LogEntity struct {
	ID int64 `gorm:"primaryKey;autoIncrement;notNull"`
	Timestamp int64 
	Request string
}