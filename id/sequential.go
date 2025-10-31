package id

import (
	"math"
	"sync"
)

// SequentialGenerator configuration.
type sequentialGeneratorConfig struct {
	first uint64 // First value in the sequence.
	last  uint64 // Last value in the sequence.
}

// SequentialGenerator option.
type SequentialGeneratorOption func(cfg *sequentialGeneratorConfig)

// Sequential number generator.
type SequentialGenerator struct {
	first uint64     // First value in the sequence.
	last  uint64     // Last value in the sequence.
	next  uint64     // Next value in the sequence.
	mutex sync.Mutex // Sync access to the next value.
}

// Optional minimum value. Defaults to 0.
func WithFirst(value uint64) SequentialGeneratorOption {
	return func(cfg *sequentialGeneratorConfig) {
		cfg.first = value
	}
}

// Optional maximum value. Defaults to math.MaxUint64.
func WithLast(value uint64) SequentialGeneratorOption {
	return func(cfg *sequentialGeneratorConfig) {
		cfg.last = value
	}
}

// Create a new SequentialGenerator.
func NewSequentialGenerator(options ...SequentialGeneratorOption) *SequentialGenerator {
	// Init default config.
	cfg := &sequentialGeneratorConfig{
		first: 0,
		last:  math.MaxUint64,
	}
	// Apply optionsto config.
	for _, option := range options {
		option(cfg)
	}

	return &SequentialGenerator{
		first: cfg.first,
		last:  cfg.last,
		next:  cfg.first,
		mutex: sync.Mutex{},
	}
}

// Get the next ID number.
func (g *SequentialGenerator) Next() uint64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	next := g.next

	if g.next < g.last {
		g.next++
	} else {
		g.next = g.first
	}

	return next
}

// Default sequential generator.
var Sequential *SequentialGenerator = NewSequentialGenerator()
