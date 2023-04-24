package service

import (
	dto "book/DTO"
	"book/entities"
	"encoding/json"
)

func parseToDb[t dto.BookDTO](dtoData t, dbData *entities.Book) error {
	data, err := json.Marshal(dtoData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dbData)
	if err != nil {
		return err
	}
	return nil
}

func parseToDTO[t dto.BookDTO ](dbData entities.Book, dtoData *t) error {
	data, err := json.Marshal(dbData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dtoData)
	if err != nil {
		return err
	}
	return nil
}
