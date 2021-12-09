package inmem

type InMemRepository struct {
	books []inmemBook
	inventory map[string]uint
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
