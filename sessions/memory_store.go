package sessions

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jrrdcnnlly/core/id"
)

// Session store held entirely in memory.
// Create with NewMemoryStore().
type MemoryStore struct {
	id       *id.RandomGenerator
	sessions map[string]*Session
	mutex    sync.Mutex
}

// Create a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		id:       id.NewRandomGenerator(),
		sessions: map[string]*Session{},
	}

	// Every hour run cleanup to remove expired sessions.
	ticker := time.NewTicker(time.Hour)
	go func() {
		for range ticker.C {
			store.Cleanup()
		}
	}()

	return store
}

// Create a new session in the store.
func (s *MemoryStore) Create() (*Session, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	session := NewSession(s.id.Next())
	s.sessions[session.id] = session
	return session, nil
}

// Retrieve a session from the store.
func (s *MemoryStore) Read(id string) (*Session, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	session, ok := s.sessions[id]
	if !ok {
		return nil, fmt.Errorf("no session with id %q", id)
	}

	if session.Expired() {
		return nil, fmt.Errorf("session %q has expired", id)
	}

	return session, nil
}

// Update a session in the store.
func (s *MemoryStore) Update(session *Session) error {
	// Pointers mean memory store sessions are always up to date.
	return nil
}

// Delete a session from the store.
func (s *MemoryStore) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.sessions, id)
	return nil
}

// Extract session ID from request and return session.
// If session does not exist create a new session.
func (s *MemoryStore) FromRequest(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(sessionCookie)
	if err == nil {
		session, err := s.Read(cookie.Value)
		if err == nil {
			return session, nil
		}
	}
	session, err := s.Create()
	if err != nil {
		return nil, fmt.Errorf("MemoryStore.FromRequest; %w", err)
	}
	return session, nil
}

// Delete expired sessions from the store.
func (s *MemoryStore) Cleanup() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for id, session := range s.sessions {
		if session.Expired() {
			delete(s.sessions, id)
		}
	}
}
