package id

import (
	"math"
	"sync"
)

type SequentialGenerator struct {
	current uint64
	mutex   sync.Mutex
}

func NewSequentialGenerator() *SequentialGenerator {
	return &SequentialGenerator{}
}

// Get the next ID number.
func (g *SequentialGenerator) Next() uint64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.current < math.MaxUint64 {
		g.current++
	} else {
		g.current = 1
	}

	return g.current
}

// Default sequential generator.
var Sequential *SequentialGenerator = &SequentialGenerator{}
