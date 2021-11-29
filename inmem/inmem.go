package inmem

type InMemRepository struct {
	books []inmemBook
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
