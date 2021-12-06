package inmem

import "bookstore/author/infrastructure/inmem"

type InMemRepository struct {
	books   []inmemBook
	authors []inmem.Author
}

func NewInMemoryRepository(dns string) *InMemRepository {
	return &InMemRepository{}
}

func (i *InMemRepository) Open() error {
	return nil
}

func (i *InMemRepository) Close() error {
	return nil
}
