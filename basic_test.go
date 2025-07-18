package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	assert.Equal(t, 66.0, Mean(53, 55, 63, 78, 81))
}
