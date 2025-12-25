package controller

import (
	"go-clean-arc/entities"
	"go-clean-arc/usecases"

	"gorm.io/gorm"
)

type BookDb struct {
	db *gorm.DB
}

func NewBookDb(db *gorm.DB) usecases.BookRepository {
	return &BookDb{db: db}
}

func (d *BookDb) Save(book entities.Book) {
	d.db.Save(&book)
}
func (d *BookDb) GetFirstBook() entities.Book {
	var book entities.Book
	d.db.First(&book)

	return book
}
