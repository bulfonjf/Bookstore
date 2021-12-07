package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidAuthorName = errors.New("first name and last name of an Author are required")
	ErrInvalidAuthorID   = errors.New("author ID must be a valid ID")
)

type Author struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}

func NewAuthor(firstName, lastName string) (Author, error) {
	if firstName == "" || lastName == "" {
		return Author{}, ErrInvalidAuthorName
	}

	return Author{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func ParseAuthorID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, ErrInvalidAuthorID
	}

	return parsedID, nil
}
