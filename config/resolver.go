package config

import (
	"fmt"
	"net/url"
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

func ConvBool(value fmt.Stringer) Resolver[bool] {
	return func(s Setting[bool]) Setting[bool] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseBool(value.String()); err == nil {
			return NewSetting(parsed)
		}
		return s
	}
}

func ConvFloat32(value fmt.Stringer) Resolver[float32] {
	return func(s Setting[float32]) Setting[float32] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseFloat(value.String(), 32); err == nil {
			return NewSetting(float32(parsed))
		}
		return s
	}
}

func ConvFloat64(value fmt.Stringer) Resolver[float64] {
	return func(s Setting[float64]) Setting[float64] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseFloat(value.String(), 64); err == nil {
			return NewSetting(float64(parsed))
		}
		return s
	}
}

func ConvInt(value fmt.Stringer) Resolver[int] {
	return func(s Setting[int]) Setting[int] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseInt(value.String(), 10, 0); err == nil {
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
		if parsed, err := strconv.ParseInt(value.String(), 10, 8); err == nil {
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
		if parsed, err := strconv.ParseInt(value.String(), 10, 16); err == nil {
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
		if parsed, err := strconv.ParseInt(value.String(), 10, 32); err == nil {
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
		if parsed, err := strconv.ParseInt(value.String(), 10, 64); err == nil {
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

func ConvUint(value fmt.Stringer) Resolver[uint] {
	return func(s Setting[uint]) Setting[uint] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseUint(value.String(), 10, 0); err == nil {
			return NewSetting(uint(parsed))
		}
		return s
	}
}

func ConvUint8(value fmt.Stringer) Resolver[uint8] {
	return func(s Setting[uint8]) Setting[uint8] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseUint(value.String(), 10, 8); err == nil {
			return NewSetting(uint8(parsed))
		}
		return s
	}
}

func ConvUint16(value fmt.Stringer) Resolver[uint16] {
	return func(s Setting[uint16]) Setting[uint16] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseUint(value.String(), 10, 16); err == nil {
			return NewSetting(uint16(parsed))
		}
		return s
	}
}

func ConvUint32(value fmt.Stringer) Resolver[uint32] {
	return func(s Setting[uint32]) Setting[uint32] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseUint(value.String(), 10, 32); err == nil {
			return NewSetting(uint32(parsed))
		}
		return s
	}
}

func ConvUint64(value fmt.Stringer) Resolver[uint64] {
	return func(s Setting[uint64]) Setting[uint64] {
		if s.Set {
			return s
		}
		if parsed, err := strconv.ParseUint(value.String(), 10, 64); err == nil {
			return NewSetting(uint64(parsed))
		}
		return s
	}
}

func ConvURL(value fmt.Stringer) Resolver[*url.URL] {
	return func(s Setting[*url.URL]) Setting[*url.URL] {
		if s.Set {
			return s
		}
		if parsed, err := url.Parse(value.String()); err == nil {
			return NewSetting(parsed)
		}
		return s
	}
}
