package application

import "bookstore/domain"

func mapToBookDto(book domain.Book) BookDTO {
	return BookDTO{
		ID: book.ID.String(),
		Title: book.Title,
	}
}
