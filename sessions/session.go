package sessions

import (
	"time"
)

type Session struct {
	id       string
	Expires  time.Time
	UserID   string
	Username string
}

// Create a new empty session with the given iD.
func NewSession(id string) *Session {
	return &Session{
		id: id,
	}
}

// Is the session expired?
func (s *Session) Expired() bool {
	return s.Expires.Before(time.Now())
}
