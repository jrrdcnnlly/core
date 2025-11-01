package session

import (
	"time"
)

type Session[T any] struct {
	id      string
	expires time.Time
	Data    T
}

// Create a new empty session with the given iD.
func NewSession[T any](id string) *Session[T] {
	return &Session[T]{
		id: id,
	}
}

// Is the session expired?
func (s *Session[T]) Expired() bool {
	return s.expires.Before(time.Now())
}
