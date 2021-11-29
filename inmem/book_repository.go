package inmem

import "bookstore/domain"

type inmemBook struct {
	id    string
	title string
}

func (i *InMemRepository) CreateBook(book domain.Book) error {
	newBook := inmemBook{
		id: book.ID.String(),
		title: book.Title,
	}
	
	i.books = append(i.books, newBook)

	return nil
}
