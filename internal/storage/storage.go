package storage

// Storage - interface for db
type Storage interface {
	Migrate() error
}
