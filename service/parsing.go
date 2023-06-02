package service

import (
	models "book/models"
	"encoding/json"
)

func parseToDb[t models.Book](dtoData t, dbData *models.Book) error {
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

func parseToDTO[t models.Book ](dbData models.Book, dtoData *t) error {
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
