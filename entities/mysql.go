package entities

type BookInterface interface {
	Create(newBook Book) error
	Get(id string) (Book, bool)
	GetAll() []Book
	Delete(id string) error
	Update(newData Book) error
}

func NewBook() BookInterface {
	return &Book{}
}

func (b Book) Create(newBook Book) error {
	db := ConnectDb()
	result := db.Model(&Book{}).Create(newBook)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b Book) Get(id string) (Book, bool) {
	db := ConnectDb()
	result, ok := db.Model(&Book{}).Get(id)
	return result.(Book), ok

}

func (b Book) GetAll() []Book {
	db := ConnectDb()
	var books []Book
	db.Model(&Book{}).Find(&books)
	return books
}

func (b Book) Delete(id string) error {
	db := ConnectDb()
	result := db.Model(&Book{}).Where("id = ?", id).Delete(&Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b Book) Update(newData Book) error {
	db := ConnectDb()
	err := db.Model(&Book{}).Where("id = ?", newData.ID).Updates(newData).Error
	if err != nil {
		return err
	}
	return nil
}
