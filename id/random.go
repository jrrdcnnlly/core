package id

import (
	"crypto/rand"
	"encoding/base64"
)

// RandomGenerator configuration.
type randomGeneratorConfig struct {
	size int
}

// RandomGenerator option.
type RandomGeneratorOption func(cfg *randomGeneratorConfig)

// Cryptographically secure random string generator.
// Create with NewRandomGenerator().
type RandomGenerator struct {
	size int
}

// Optional number of random bytes.
func WithSize(size int) RandomGeneratorOption {
	return func(cfg *randomGeneratorConfig) {
		cfg.size = size
	}
}

// Create a new RandomGenerator
func NewRandomGenerator(options ...RandomGeneratorOption) *RandomGenerator {
	// Init default config.
	cfg := &randomGeneratorConfig{
		size: 32,
	}
	// Apply options to config.
	for _, option := range options {
		option(cfg)
	}

	return &RandomGenerator{
		size: cfg.size,
	}
}

// Generate an ID string.
func (g *RandomGenerator) Next() string {
	buffer := make([]byte, g.size)
	rand.Read(buffer) // Ignore returned values.
	return base64.StdEncoding.EncodeToString(buffer)
}
