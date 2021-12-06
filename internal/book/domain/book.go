package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidBookTitle = errors.New("Invalid title")
)

type Book struct {
	ID    uuid.UUID
	Title string
}

func NewBook(title string) (Book, error) {
	if title == "" {
		return Book{}, ErrInvalidBookTitle
	}

	return Book{
		ID:    uuid.New(),
		Title: title,
	}, nil
}
