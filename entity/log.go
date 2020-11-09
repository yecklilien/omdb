package entity

//Log entity
type Log struct {
	ID int64 `gorm:"primaryKey;autoIncrement;notNull"`
	Timestamp int64 
	Request string
}