package entities

type Book struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"type:varchar(100)"`
	Author      string `gorm:"type:varchar(100)"`
}
