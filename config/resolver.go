package config

import (
	"fmt"
	"strconv"
)

type Resolver[T any] func(s Setting[T]) Setting[T]

func Fallback[T any](value T) Resolver[T] {
	return func(s Setting[T]) Setting[T] {
		if s.Set {
			return s
		}
		return NewSetting(value)
	}
}

func PanicIfUnset[T any]() Resolver[T] {
	return func(s Setting[T]) Setting[T] {
		if s.Set {
			return s
		}
		panic("config value is not set")
	}
}

func ConvInt(value fmt.Stringer) Resolver[int] {
	return func(s Setting[int]) Setting[int] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 0); err != nil {
			return NewSetting(int(parsed))
		}
		return s
	}
}

func ConvInt8(value fmt.Stringer) Resolver[int8] {
	return func(s Setting[int8]) Setting[int8] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 8); err != nil {
			return NewSetting(int8(parsed))
		}
		return s
	}
}

func ConvInt16(value fmt.Stringer) Resolver[int16] {
	return func(s Setting[int16]) Setting[int16] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 16); err != nil {
			return NewSetting(int16(parsed))
		}
		return s
	}
}

func ConvInt32(value fmt.Stringer) Resolver[int32] {
	return func(s Setting[int32]) Setting[int32] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 32); err != nil {
			return NewSetting(int32(parsed))
		}
		return s
	}
}

func ConvInt64(value fmt.Stringer) Resolver[int64] {
	return func(s Setting[int64]) Setting[int64] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 64); err != nil {
			return NewSetting(int64(parsed))
		}
		return s
	}
}

func ConvString(value fmt.Stringer, allowEmpty bool) Resolver[string] {
	return func(s Setting[string]) Setting[string] {
		if s.Set {
			return s
		}
		if !allowEmpty && value.String() == "" {
			return s
		}
		return NewSetting(value.String())
	}
}
