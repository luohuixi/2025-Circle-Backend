package models
type Circle struct {
    Id int `gorm:"primaryKey;autoIncrement"`
	Name string
	Imageurl string
	Discription string `gorm:"type:text"`
	Userid int
	Status string
}
type FollowCircle struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Circleid int
	Userid int
}