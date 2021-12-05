package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidTitle = errors.New("Invalid title")
)

type Book struct {
	ID    uuid.UUID
	Title string
}

func NewBook(title string) (Book, error) {
	if title == "" {
		return Book{}, ErrInvalidTitle
	}

	return Book{
		ID:    uuid.New(),
		Title: title,
	}, nil
}
