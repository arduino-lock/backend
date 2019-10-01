package golockserver

import "github.com/boltdb/bolt"

// DatabaseService is an interface with database-specific methods
type DatabaseService interface {
	Setup(db *bolt.DB) error
	DatabasePrint(development bool) error
}
