package inmem

type InMemRepository struct {
	DNS        string
	tables     map[string][]interface{}
	migrations []func(db *InMemRepository) error
}

func NewInMemRepository(dns string) *InMemRepository {
	return &InMemRepository{
		DNS:        dns,
		tables:     make(map[string][]interface{}),
		migrations: make([]func(db *InMemRepository) error, 0),
	}
}

func (i *InMemRepository) AddTable(name string, collection []interface{}) {
	i.tables[name] = collection
}

func (i *InMemRepository) AddMigration(migration func(db *InMemRepository) error) {
	i.migrations = append(i.migrations, migration)
}

func (i *InMemRepository) Open() error {
	for _, m := range i.migrations {
		if err := m(i); err != nil {
			return err
		}
	}

	return nil
}

func (i *InMemRepository) Close() error {
	return nil
}
