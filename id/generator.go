package id

// Define an interface for returning ID values.
type IDGenerator[T any] interface {
	Next() T
}
