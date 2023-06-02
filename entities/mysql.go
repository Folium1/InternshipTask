package entities

import (
	"book/config"
	"book/models"
)

type BookInterface interface {
	Create(newBook models.Book) (int, error)
	Get(id string) (models.Book, error)
	GetAll() []models.Book
	Delete(id string) error
	Update(newData models.Book) error
}

type Book struct {
	models.Book
}

func NewBook() BookInterface {
	return &Book{}
}

func (b Book) Create(newBook models.Book) (int, error) {
	db := config.ConnectDb()
	err := db.Create(&newBook).Error
	if err != nil {
		return -1, err
	}
	return newBook.ID, nil
}

func (b Book) Get(id string) (models.Book, error) {
	db := config.ConnectDb()
	var book models.Book
	err := db.Where("id = ?", id).First(&book).Error
	return book, err
}

func (b Book) GetAll() []models.Book {
	db := config.ConnectDb()
	var books []models.Book
	db.Find(&books)
	return books
}

func (b Book) Delete(id string) error {
	db := config.ConnectDb()
	result := db.Where("id = ?", id).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b Book) Update(newData models.Book) error {
	db := config.ConnectDb()
	err := db.Where("id = ?", newData.ID).Updates(newData).Error
	if err != nil {
		return err
	}
	return nil
}
