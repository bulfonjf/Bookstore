package application

import (
	"bookstore/domain"
	"errors"
)

func NewInventoryService(repository InventoryRepository, bookRepository BookRepository) *InventoryService {
	return &InventoryService{
		repository: repository,
		bookRepository: bookRepository,
	}
}

type InventoryService struct {
	repository InventoryRepository
	bookRepository BookRepository
}

func (i *InventoryService) AddBook(bookDTO BookDTO) error {
	book, err := mapToBook(bookDTO)
	if err != nil {
		return Error{Code: EINTERNAL, Message: err.Error()}
	}

	if err = i.repository.AddBook(book, 1); err != nil {
		return Error{Code: EINTERNAL, Message: err.Error()}
	}

	return nil
}

func (i *InventoryService) GetInventory() ([]InventoryDetailDTO, error) {
	inventoryDTO := make([]InventoryDetailDTO, 0)

	inventory, err := i.repository.GetInventory()
	if err != nil {
		return []InventoryDetailDTO{}, Error{Code: EINTERNAL, Message: err.Error()}
	}

	for k, v := range inventory{
		bookID, err := ParseBookID(k)
		if err != nil {
			return []InventoryDetailDTO{}, Error{Code: EINTERNAL, Message: err.Error()}
		}

		book, err := i.bookRepository.GetBookByID(bookID)
		if err != nil && errors.Is(err, ErrNotFound) {
			// todo delete book from inventory
			continue
		} else if err != nil {
			return []InventoryDetailDTO{}, Error{Code: EINTERNAL, Message: err.Error()}
		}

		inventoryDetail := domain.NewInventoryDetail(book, v)

		inventoryDTO = append(inventoryDTO, mapToInventoryDetailDTO(inventoryDetail))
	}

	return inventoryDTO, nil
}
