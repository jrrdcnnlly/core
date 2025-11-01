package session

import (
	"context"
	"errors"
)

// Used as a unique context key.
type sessionKey struct{}

// Retrieve a session from a context.
func FromContext[T any](ctx context.Context) (*Session[T], error) {
	value := ctx.Value(sessionKey{})
	if value == nil {
		return nil, errors.New("no session in context")
	}
	session, ok := value.(*Session[T])
	if !ok {
		return nil, errors.New("invalid session in context")
	}
	return session, nil
}
