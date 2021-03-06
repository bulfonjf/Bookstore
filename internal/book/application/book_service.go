package application

import (
	"bookstore/domain"
	"bookstore/internal/apperror"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrCreateBook    = apperror.Error{Code: apperror.EINTERNAL, Message: "the book can't be created in the repository"}
	ErrNotFound      = apperror.Error{Code: apperror.ENOTFOUND, Message: "the book doesn't exist in the repository"}
	ErrInvalidBookID = apperror.Error{Code: apperror.EINVALID, Message: "Book id must be a valid uuid"}
)

type BookService struct {
	repository BookRepository
}

func NewBookService(repository BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (bs *BookService) CreateBook(createBookDTO CreateBookDTO) (BookDTO, error) {
	book, err := domain.NewBook(createBookDTO.Title)
	if err != nil && errors.Is(err, domain.ErrInvalidBookTitle) {
		return BookDTO{}, Error{Code: EINVALID, Message: err.Error()}
	}

	if err := bs.repository.CreateBook(book); err != nil {
		return BookDTO{}, err
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) GetBooks() ([]BookDTO, error) {
	var booksDTO []BookDTO
	var books []domain.Book
	books, err := bs.repository.GetBooks()
	if err != nil {
		return booksDTO, err
	}

	for _, b := range books {
		booksDTO = append(booksDTO, mapToBookDTO(b))
	}

	return booksDTO, nil
}

func (bs *BookService) GetBookByID(id string) (BookDTO, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return BookDTO{}, ErrInvalidBookID
	}

	book, err := bs.repository.GetBookByID(bookID)
	if err != nil && errors.Is(err, ErrNotFound) {
		return BookDTO{}, err
	} else if err != nil {
		return BookDTO{}, fmt.Errorf("Getting book by id from repository: Error: %w", err)
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) UpdateBook(updateBookDTO UpdateBookDTO) (BookDTO, error) {
	book, err := mapToBook(updateBookDTO)
	if err != nil {
		return BookDTO{}, err
	}

	err = bs.repository.UpdateBook(book)
	if err != nil {
		return BookDTO{}, fmt.Errorf("Book Service: can't update book. Error: %w", err)
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) DeleteBook(bookID string) error {
	id, err := ParseBookID(bookID)
	if err != nil {
		return err
	}

	err = bs.repository.DeleteBook(id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("Book Service: can't delete book. Error: %w", err)
	}

	return nil
}

func ParseBookID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, ErrInvalidBookID
	}

	return parsedID, nil
}
