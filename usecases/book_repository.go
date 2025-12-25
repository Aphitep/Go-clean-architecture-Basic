package usecases

import "go-clean-arc/entities"

type BookRepository interface {
	Save(book entities.Book)
	GetFirstBook() entities.Book
}
