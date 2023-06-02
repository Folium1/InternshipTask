package models

type Book struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Title  string `gorm:"type:varchar(100)" json:"title"`
	Author string `gorm:"type:varchar(100)" json:"author"`
}
