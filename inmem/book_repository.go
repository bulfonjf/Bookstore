package inmem

import (
	"bookstore/domain"
	"fmt"

	"github.com/google/uuid"
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
		id, err := uuid.FromBytes([]byte(b.id))
		if err != nil {
			return []domain.Book{}, fmt.Errorf("InMemRepository: can't convert string ID to uuid, %w", err)
		}

		books = append(books, domain.Book{ID: id, Title: b.title})
	}

	return books, nil
}
