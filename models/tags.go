package models

type Tags struct {
	Uuid       int    `gorm:"type:int;primary_key"`
	Name_first string `gorm:"type:varchar(250)"`
	Name_last  string `gorm:"type:varchar(250)"`
	Department string `grom:"type:varchar(150);foreign_key"`
	Salary     int    `gorm:"type:int;foreign_key"`
}
