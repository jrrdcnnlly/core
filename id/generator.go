package id

type IDGenerator[T any] interface {
	Next() T
}
