package application

import "bookstore/internal/author/domain"

func mapToAuthorDTO(author domain.Author) AuthorDTO {
	return AuthorDTO{
		ID:        author.ID.String(),
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}

func mapToAuthor(updateAuthorDTO UpdateAuthorDTO) (domain.Author, error) {
	parsedID, err := domain.ParseAuthorID(updateAuthorDTO.ID)
	if err != nil {
		return domain.Author{}, err
	}

	return domain.Author{
		ID:        parsedID,
		FirstName: updateAuthorDTO.FirstName,
		LastName:  updateAuthorDTO.LastName,
	}, nil
}
