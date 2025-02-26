package models
type SearchHistory struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	SearchKey string `gorm:"index"`
	Userid int
}