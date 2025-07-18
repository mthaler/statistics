package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	assert.Equal(t, 66.0, Mean(53, 55, 63, 78, 81))
}

func TestMedian(t *testing.T) {
	assert.Equal(t, 63, Median(53, 55, 63, 78, 81))
}

func TestVariance(t *testing.T) {
	assert.Equal(t, 1048.0, Variance(78, 63, 53, 91, 55))
}
