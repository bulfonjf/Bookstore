package application

import (
	"bookstore/internal/apperror"
	"bookstore/internal/author/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrNotFound = apperror.Error{Code: apperror.ENOTFOUND, Message: "the author doesn't exist in the repository"}
)

type AuthorService struct {
	repository AuthorRepository
}

func NewAuthorService(repository AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}

func (as *AuthorService) CreateAuthor(createAuthorDTO CreateAuthorDTO) (AuthorDTO, error) {
	author, err := domain.NewAuthor(createAuthorDTO.FirstName, createAuthorDTO.LastName)
	if err != nil && errors.Is(err, domain.ErrInvalidAuthorName) {
		return AuthorDTO{}, apperror.Error{Code: apperror.EINVALID, Message: err.Error()}
	}

	if err := as.repository.CreateAuthor(author); err != nil {
		return AuthorDTO{}, err
	}

	return mapToAuthorDTO(author), nil

}

func (as *AuthorService) GetAuthors() ([]AuthorDTO, error) {
	var authorsDTO []AuthorDTO
	var authors []domain.Author
	authors, err := as.repository.GetAuthors()
	if err != nil {
		return []AuthorDTO{}, err
	}

	for _, a := range authors {
		authorsDTO = append(authorsDTO, mapToAuthorDTO(a))
	}

	return authorsDTO, nil
}

func (as *AuthorService) GetAuthorByID(id string) (AuthorDTO, error) {
	authorID, err := uuid.Parse(id)
	if err != nil {
		return AuthorDTO{}, domain.ErrInvalidAuthorID
	}

	author, err := as.repository.GetAuthorByID(authorID)
	if err != nil && errors.Is(err, ErrNotFound) {
		return AuthorDTO{}, err
	} else if err != nil {
		return AuthorDTO{}, fmt.Errorf("Getting author by id from repository, Error: %w", err)
	}

	return mapToAuthorDTO(author), nil
}

func (as *AuthorService) UpdateAuthor(updateAuthorDTO UpdateAuthorDTO) (AuthorDTO, error) {
	author, err := mapToAuthor(updateAuthorDTO)
	if err != nil {
		return AuthorDTO{}, err
	}

	err = as.repository.UpdateAuthor(author)
	if err != nil {
		return AuthorDTO, fmt.Errorf("Author Service: can't update author. Error %w", err)
	}

	return mapToAuthorDTO(author), nil
}

func (as *AuthorService) DeleteAuthor(authorID string) error {
	id, err := domain.ParseAuthorID(authorID)
	if err != nil {
		return apperror.Error{Code: apperror.EINVALID, Message: err.Error()}
	}

	err = as.repository.DeleteAuthor(id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("Author Service: can't delete author. Error %w", err)
	}

	return nil
}
