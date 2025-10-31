package id

import (
	"math"
	"testing"
)

func TestNewSequentialGenerator(t *testing.T) {
	type Test struct {
		name        string
		expectFirst uint64
		expectLast  uint64
	}

	type TestBatch struct {
		gen   *SequentialGenerator
		tests []Test
	}

	batches := []TestBatch{
		{
			gen: NewSequentialGenerator(),
			tests: []Test{
				{".first should be 0, .last should be math.MaxUint64", 0, math.MaxUint64},
			},
		},
		{
			gen: NewSequentialGenerator(WithFirst(10), WithLast(20)),
			tests: []Test{
				{".first should be 10, .last should be 20", 10, 20},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				first := job.gen.first
				last := job.gen.last
				if test.expectFirst != first {
					t.Errorf("SequentialGenerator.first = %d; expected %d", first, test.expectFirst)
				}
				if test.expectLast != last {
					t.Errorf("SequentialGenerator.last = %d; expected %d", last, test.expectLast)
				}
			})
		}
	}
}

func TestSequentialGenerator_Next(t *testing.T) {
	type Test struct {
		name   string
		expect uint64
	}

	type TestBatch struct {
		gen   *SequentialGenerator
		tests []Test
	}

	batches := []TestBatch{
		{
			gen: NewSequentialGenerator(),
			tests: []Test{
				{"First call should be 0", 0},
				{"Second call should be 1", 1},
				{"Third call should be 2", 2},
				{"Fourth call should be 3", 3},
			},
		},
		{
			gen: NewSequentialGenerator(WithFirst(10), WithLast(12)),
			tests: []Test{
				{"First call should be 10", 10},
				{"Second call should be 11", 11},
				{"Third call should be 12", 12},
				{"Fourth call should be 130", 10},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				result := job.gen.Next()
				if test.expect != result {
					t.Errorf("SequentialGenerator.Next() = %d; expected %d", result, test.expect)
				}
			})
		}
	}
}

func BenchmarkSequentialGenerator_Next(b *testing.B) {
	gen := NewSequentialGenerator()

	for b.Loop() {
		gen.Next()
	}
}
