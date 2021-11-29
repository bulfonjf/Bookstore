package application

import "bookstore/domain"

type BookRepository interface {
	CreateBook(book domain.Book) error
}

