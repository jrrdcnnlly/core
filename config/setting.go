package config

// Application setting.
type Setting[T any] struct {
	Value T
	Set   bool
}

// Create a new application setting with the given value.
func NewSetting[T any](value T) Setting[T] {
	return Setting[T]{
		Value: value,
		Set:   true,
	}
}

// Resolve the setting value from one or more resolvers.
func (s Setting[T]) Resolve(resolvers ...Resolver[T]) T {
	for _, resolver := range resolvers {
		s = resolver(s)
	}
	return s.Value
}
