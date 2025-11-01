package session

import "net/http"

// Define an interface for session caches.
type SessionStore[T any] interface {
	Create() (*Session[T], error)
	Read(id string) (*Session[T], error)
	Update(session *Session[T]) error
	Delete(id string) error
	FromRequest(r *http.Request) (*Session[T], error)
}
