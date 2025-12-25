package usecases

import "go-clean-arc/entities"

type BookUseCase interface {
	CreateBook(book entities.Book)
	FirstBook() entities.Book
}

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookUseCase {
	return &BookService{repo: repo}
}

func (r *BookService) CreateBook(book entities.Book) {
	r.repo.Save(book)
}

func (r *BookService) FirstBook() entities.Book {
	return r.repo.GetFirstBook()
}
