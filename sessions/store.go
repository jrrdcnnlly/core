package sessions

import "net/http"

// Define an interface for session caches.
type SessionStore interface {
	Create() (*Session, error)
	Read(id string) (*Session, error)
	Update(session *Session) error
	Delete(id string) error
	FromRequest(r *http.Request) (*Session, error)
}
