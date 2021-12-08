package sqlite

import (
	"bookstore/internal/author/application"
	"bookstore/internal/author/domain"
	"bookstore/internal/platform/db/sqlite"
	"context"
	"fmt"

	"github.com/google/uuid"
)

const (
	authorTable     = "Authors"
	firstNameColumn = "FirstName"
	lastNameColumn  = "LastName"
)

var (
	insertAuthorQuery = fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES (?,?)", authorTable, firstNameColumn, lastNameColumn)
)

type Author struct {
	id        string
	firstName string
	lastName  string
}

type SQLiteAuthorRepository struct {
	db *sqlite.DB
}

func NewSQLiteAuthorRepository(db *sqlite.DB) *SQLiteAuthorRepository {
	return &SQLiteAuthorRepository{
		db: db,
	}
}

func Migration(db *sqlite.DB) error {

}

func (s *SQLiteAuthorRepository) CreateAuthor(ctx context.Context, author domain.Author) error {
	tx, err := s.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, insertAuthorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, author.FirstName, author.LastName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *SQLiteAuthorRepository) GetAuthors() ([]domain.Author, error) {
	tx, err := s.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, getAuthorsQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var dbAuthors Author

	err = stmt.GetContext(ctx, &dbAuthors)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	var authors []domain.Author

	for _, a := range dbAuthors {
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

func (s *SQLiteAuthorRepository) GetAuthorByID(id uuid.UUID) (domain.Author, error) {
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

func (s *SQLiteAuthorRepository) UpdateAuthor(author domain.Author) error {
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

func (s *SQLiteAuthorRepository) DeleteAuthor(id uuid.UUID) error {
	_, err := i.GetAuthorByID(id)
	if err != nil {
		return err
	}

	authorIndex := i.getAuthorIndex(id)
	i.authors = deleteAuthor(i.authors, authorIndex)

	return nil
}

func (s *SQLiteAuthorRepository) getAuthorIndex(id uuid.UUID) int {
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
