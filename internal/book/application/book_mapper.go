package application

import "bookstore/domain"

func mapToBookDTO(book domain.Book) BookDTO {
	return BookDTO{
		ID:    book.ID.String(),
		Title: book.Title,
	}
}

func mapToBook(book UpdateBookDTO) (domain.Book, error) {
	id, err := ParseBookID(book.ID)
	if err != nil {
		return domain.Book{}, err
	}

	return domain.Book{
		ID:    id,
		Title: book.Title,
	}, nil
}
