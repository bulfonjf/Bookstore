package domain

import "github.com/google/uuid"

type Book struct {
	ID    uuid.UUID
	Title string
}

func NewBook(title string) Book {
	return Book{
		ID: uuid.New(),
		Title: title,
	}
}