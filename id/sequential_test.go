package id

import (
	"fmt"
	"testing"
)

func TestSequentialGenerator_Next(t *testing.T) {

	t.Run("NewSequentialGenerator()", func(t *testing.T) {
		t.Parallel()

		var (
			gen      *SequentialGenerator
			result   uint64
			expected uint64
		)

		gen = NewSequentialGenerator()

		result = gen.Next()
		expected = 0
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}

		result = gen.Next()
		expected = 1
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}

		result = gen.Next()
		expected = 2
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}
	})

	var (
		first uint64 = 10
		last  uint64 = 11
	)

	t.Run(fmt.Sprintf("NewSequentialGenerator(%d, %d)", first, last), func(t *testing.T) {
		t.Parallel()

		var (
			gen      *SequentialGenerator
			result   uint64
			expected uint64
		)

		gen = NewSequentialGenerator(WithFirst(first), WithLast(last))

		result = gen.Next()
		expected = 10
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}

		result = gen.Next()
		expected = 11
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}

		result = gen.Next()
		expected = 10
		if result != expected {
			t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, expected)
		}
	})

}

func BenchmarkSequentialGenerator_Next(b *testing.B) {
	gen := NewSequentialGenerator()

	for i := 0; i < b.N; i++ {
		gen.Next()
	}
}
