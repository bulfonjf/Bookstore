package inmem

import (
	"bookstore/application"
	"bookstore/author/domain"

	"github.com/google/uuid"
)

type Author struct {
	id        string
	firstName string
	lastName  string
}

type InMemAuthorRepository struct {
	authors []Author
}

func (i *InMemAuthorRepository) CreateAuthor(author domain.Author) error {
	newAuthor := Author{
		id:        author.ID.String(),
		firstName: author.FirstName,
		lastName:  author.LastName,
	}

	i.authors = append(i.authors, newAuthor)

	return nil
}

func (i *InMemAuthorRepository) GetAuthors() ([]domain.Author, error) {
	var authors []domain.Author

	for _, a := range i.authors {
		parsedID, err := domain.ParseAuthorID(a.id)
		if err != nil {
			return []domain.Author{}, err
		}

		authors = append(authors, domain.Author{
			ID:        parsedID,
			FirstName: a.firstName,
			LastName:  a.lastName,
		})
	}

	return authors, nil
}

func (i *InMemAuthorRepository) GetAuthorByID(id uuid.UUID) (domain.Author, error) {
	authorFound := domain.Author{}
	authorIndex := i.getAuthorIndex(id)
	if authorIndex < 0 {
		return domain.Author{}, application.ErrNotFound
	}

	a := i.authors[authorIndex]
	parsedID, err := domain.ParseAuthorID(a.id)
	if err != nil {
		return domain.Author{}, err
	}

	authorFound = domain.Author{
		ID:        parsedID,
		FirstName: a.firstName,
		LastName:  a.lastName,
	}

	return authorFound, nil
}

func (i *InMemAuthorRepository) UpdateAuthor(author domain.Author) error {
	authorIndex := i.getAuthorIndex(author.ID)
	if authorIndex < 0 {
		return i.CreateAuthor(author)
	} else {
		i.authors[authorIndex] = Author{
			id:        author.ID.String(),
			firstName: author.FirstName,
			lastName:  author.LastName,
		}

		return nil
	}
}

func (i *InMemAuthorRepository) DeleteAuthor(id uuid.UUID) error {
	_, err := i.GetAuthorByID(id)
	if err != nil {
		return err
	}

	authorIndex := i.getAuthorIndex(id)
	i.authors = deleteAuthor(i.authors, authorIndex)

	return nil
}

func (i *InMemAuthorRepository) getAuthorIndex(id uuid.UUID) int {
	for index, b := range i.authors {
		if id.String() == b.id {
			return index
		}

	}

	return -1
}

func deleteAuthor(authors []Author, indexToRemove int) []Author {
	currentLength := len(authors)

	if currentLength == 0 {
		return authors
	}

	lastItem := currentLength - 1

	switch true {
	case indexToRemove == lastItem:
		authors = authors[:lastItem]
	case indexToRemove > 0 && indexToRemove < lastItem:
		authors[indexToRemove] = authors[lastItem]
		authors = authors[:lastItem]
	case indexToRemove == 0:
		authors = authors[1:]
	}

	return authors
}
