package service

import (
	"book/config"
	"book/entities"
	models "book/models"
	"errors"
	"fmt"
)

type Service struct {
	db entities.BookInterface
}

func New() Service {
	return Service{db: entities.NewBook()}
}

func (s Service) Create(newBook models.Book) (int, error) {
	var dbBook models.Book
	err := parseToDb(newBook, &dbBook)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't parse data to db, err: %v", err.Error()))
		return -1, errors.New("Couldn't parse data to db")
	}
	id, err := s.db.Create(dbBook)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't create new book,error: %v", err.Error()))
		return -1, errors.New("Couldn't create new book")
	}
	return id, nil
}

func (s Service) Get(id string) (models.Book, error) {
	book, err := s.db.Get(id)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't get the book, err: %v", err.Error()))
		return models.Book{}, errors.New("No book with that id")
	}
	var bookDTO models.Book
	err = parseToDTO(book, &bookDTO)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't parse data to dto, err: %v", err.Error()))
		return models.Book{}, errors.New("Couldn't parse data to dto")
	}
	return bookDTO, nil
}

func (s Service) Delete(id string) error {
	err := s.db.Delete(id)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't delete the book, err: %v", err.Error()))
		return errors.New("Couldn't delete the book")
	}
	return nil
}

func (s Service) GetAll() ([]models.Book, error) {
	var booksDTO []models.Book
	dbBooks := s.db.GetAll()
	for _, i := range dbBooks {
		var bookDTO models.Book
		err := parseToDTO(i, &bookDTO)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't parse data to dto, err: %v", err.Error()))
			return nil, errors.New("Couldn't parse data to dto")
		}
		booksDTO = append(booksDTO, bookDTO)
	}
	return booksDTO, nil
}

func (s Service) Update(newData models.Book) error {
	var dbData models.Book
	err := parseToDb(newData, &dbData)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't parse data to db, err: %v", err.Error()))
		return errors.New("Couldn't parse data to db")
	}
	err = s.db.Update(dbData)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't update book's data, err: %v", err.Error()))
		return errors.New("Couldn't update book's data")
	}
	return nil
}
