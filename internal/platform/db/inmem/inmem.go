package inmem

type InMemRepository struct {
	DNS string
}

func (i *InMemRepository) Open() error {
	return nil
}

func (i *InMemRepository) Close() error {
	return nil
}
