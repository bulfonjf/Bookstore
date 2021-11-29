package application

import (
	"bookstore/domain"
	"errors"
)

var (
	ErrCreateBook = errors.New("the book can't be created in the repository")
)

func NewBookService(repository BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

type BookService struct {
	repository BookRepository
}

func (bs *BookService) CreateBook(createBookDTO CreateBookDTO) (BookDTO, error) {
	book := domain.NewBook(createBookDTO.Title)

	if err := bs.repository.CreateBook(book); err != nil {
		return BookDTO{}, err
	}

	return mapToBookDto(book), nil
}

func (bs *BookService) GetBooks() ([]BookDTO, error) {
	var booksDTO []BookDTO
	var books []domain.Book
	books, err := bs.repository.GetBooks()
	if  err != nil {
		return booksDTO, err
	}

	for _, b := range books {
		booksDTO = append(booksDTO, mapToBookDto(b))
	}

	return booksDTO, nil
}
