package id

import (
	"encoding/base64"
	"testing"
)

func TestNewRandomGenerator(t *testing.T) {
	expectedSize := 32

	gen := NewRandomGenerator()

	if gen.size != expectedSize {
		t.Errorf("NewRandomGenerator().size = %d; expected %d", gen.size, expectedSize)
	}

}

func TestRandomGenerator_Next(t *testing.T) {
	t.Run("NewRandomGenerator()", func(t *testing.T) {
		expectedSize := 32

		t.Parallel()

		gen := NewRandomGenerator()

		id := gen.Next()
		decoded, err := base64.StdEncoding.DecodeString(id)
		if err != nil {
			t.Error(err)
		}
		if len(decoded) != expectedSize {
			t.Errorf("len(decoded) = %d; expected %d", len(decoded), expectedSize)
		}
	})

	t.Run("NewRandomGenerator()", func(t *testing.T) {
		expectedSize := 128

		t.Parallel()

		gen := NewRandomGenerator(WithSize(expectedSize))

		id := gen.Next()
		decoded, err := base64.StdEncoding.DecodeString(id)
		if err != nil {
			t.Error(err)
		}
		if len(decoded) != expectedSize {
			t.Errorf("len(decoded) = %d; expected %d", len(decoded), expectedSize)
		}
	})
}
