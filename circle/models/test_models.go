package models
import "time"
type Test struct {
    Testid int `gorm:"primaryKey;autoIncrement"`	
	Testname string
	Userid int
	Discription string `gorm:"type:text"`
	Circle string
	Good int
	Status string
	Createtime time.Time `gorm:"autoCreateTime"`
	Imageurl string
}
type TestQuestion struct {
	Testid int
	Questionid int `gorm:"primaryKey;autoIncrement"`	
	Content string `gorm:"type:text"`
	Difficulty string
	Answer string
	Variety string
	Imageurl string
	Explain string `gorm:"type:text"`
}
type TestOption struct {
	Optionid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Practiceid int
	Option string
}
type Top struct {
	Topid int `gorm:"primaryKey;autoIncrement"`
	Userid int
	Correctnum int
	Time int
	Testid int
}
type TestComment struct {
	Commentid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Testid int
	Userid int
	Createtime time.Time `gorm:"autoCreateTime"`
	Good int
}
type Testhistory struct {
	Testhisrotyid int `gorm:"primaryKey;autoIncrement"`
	Userid int
	Testid int
}