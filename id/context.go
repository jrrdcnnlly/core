package id

import (
	"context"
)

type idKey struct{}

func FromContext[T any](ctx context.Context) (value T, ok bool) {
	raw := ctx.Value(idKey{})
	value, ok = raw.(T)
	return
}
