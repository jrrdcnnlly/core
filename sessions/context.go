package sessions

import (
	"context"
	"errors"
)

// Used as a unique context key.
type sessionKey struct{}

// Retrieve a session from a context.
func FromContext(ctx context.Context) (*Session, error) {
	value := ctx.Value(sessionKey{})
	if value == nil {
		return nil, errors.New("no session in context")
	}
	session, ok := value.(*Session)
	if !ok {
		return nil, errors.New("invalid session in context")
	}
	return session, nil
}
