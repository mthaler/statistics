package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponential(t *testing.T) {
	l := 5730.0
	x := 1.0 / 8267.0
	assert.Equal(t, 0.4999850160033138, 1-Exponential(x, l)/l)
}
