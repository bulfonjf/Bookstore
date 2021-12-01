package inmem

import (
	"bookstore/application"
	"bookstore/domain"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("book not found in repository")
)

type inmemBook struct {
	id    string
	title string
}

func (i *InMemRepository) CreateBook(book domain.Book) error {
	newBook := inmemBook{
		id:    book.ID.String(),
		title: book.Title,
	}

	i.books = append(i.books, newBook)

	return nil
}

func (i *InMemRepository) GetBooks() ([]domain.Book, error) {
	var books []domain.Book

	for _, b := range i.books {
		parsedID, err := application.ParseBookID(b.id)
		if err != nil {
			return []domain.Book{}, err
		}

		books = append(books, domain.Book{ID: parsedID, Title: b.title})
	}

	return books, nil
}

func (i *InMemRepository) GetBookByID(id uuid.UUID) (domain.Book, error) {
	bookFound := domain.Book{}
	bookIndex := i.getBookIndex(id)
	if bookIndex > 0 {
		b := i.books[bookIndex]
		parsedID, err := application.ParseBookID(b.id)
		if err != nil {
			return domain.Book{}, err
		}

		bookFound = domain.Book{ID: parsedID, Title: b.title}
	}

	if domain.BooksAreEqual(bookFound, domain.Book{}) {
		return domain.Book{}, ErrNotFound
	}

	return bookFound, nil
}

func (i *InMemRepository) UpdateBook(book domain.Book) error {
	bookIndex := i.getBookIndex(book.ID)
	if bookIndex < 0 {
		return i.CreateBook(book)
	} else {
		i.books[bookIndex] = inmemBook{id: book.ID.String(), title: book.Title}
		return nil
	}
}

func (i *InMemRepository) getBookIndex(id uuid.UUID) int {
	for index, b := range i.books {
		if b.id == id.String() {
			return index, nil
		}

	}

	return 0, nil
}
