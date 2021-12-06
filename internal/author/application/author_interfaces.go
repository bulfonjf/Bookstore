package application

import (
	"bookstore/internal/author/domain"

	"github.com/google/uuid"
)

type AuthorRepository interface {
	CreateAuthor(author domain.Author) error
	DeleteAuthor(id uuid.UUID) error
	GetAuthors() ([]domain.Author, error)
	GetAuthorByID(id uuid.UUID) (domain.Author, error)
	UpdateAuthor(author domain.Author) error
}
