package config

type Setting[T any] struct {
	Value T
	Set   bool
}

func NewSetting[T any](value T) Setting[T] {
	return Setting[T]{
		Value: value,
		Set:   true,
	}
}

func (s Setting[T]) Resolve(resolvers ...Resolver[T]) T {
	for _, resolver := range resolvers {
		s = resolver(s)
	}
	return s.Value
}
