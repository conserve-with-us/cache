package cache

import (
	"time"
)

// Cache contains state around initializing a connection to a cache layer
type Cache interface {
	Open(uri string) (Conn, error)
}

// Conn is a connection to a cache layer or data store
type Conn interface {
	Write(k, v []byte) error
	WriteTTL(k, v []byte, ttl time.Duration) error
	Read(k []byte) ([]byte, error)
	Close() error
	Stats() (map[string]interface{}, error)
}
