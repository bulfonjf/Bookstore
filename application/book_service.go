package application

import (
	"bookstore/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
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
	if err != nil {
		return booksDTO, err
	}

	for _, b := range books {
		booksDTO = append(booksDTO, mapToBookDto(b))
	}

	return booksDTO, nil
}

func (bs *BookService) GetBookByID(id string) (BookDTO, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return BookDTO{}, fmt.Errorf("Getting book by id: Invalid id. Error: %w", err)
	}

	book, err := bs.repository.GetBookByID(bookID)
	if err != nil {
		return BookDTO{}, fmt.Errorf("Getting book by id from repository: Error: %w", err)
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) UpdateBook(updateBookDTO UpdateBookDTO) (BookDTO, error) {
	book, err := mapToBook(updateBookDTO)
	if err != nil {
		return BookDTO{}, err
	}

	updatedBook, err := bs.repository.UpdateBook(book)
	if err != nil {
		return BookDTO{}, fmt.Errorf("Book Service: can't update book. Error: %w", err)
	}

	return mapToBookDTO(updatedBook), nil
}

func ParseBookID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Book Service: can't convert string ID to uuid, %w", err)
	}

	return parsedID, nil
}
