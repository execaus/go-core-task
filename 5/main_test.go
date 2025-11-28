package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectInts(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "normal case with intersection",
			a:        []int{1, 2, 3, 4},
			b:        []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "no intersection",
			a:        []int{1, 2},
			b:        []int{3, 4},
			expected: []int{},
		},
		{
			name:     "first slice empty",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "second slice empty",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "both slices empty",
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "slices with duplicates",
			a:        []int{1, 2, 2, 3, 3, 3},
			b:        []int{2, 3, 3, 4},
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := IntersectInts(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "IntersectInts(%v, %v) result", tt.a, tt.b)
			expectedOk := len(tt.expected) > 0
			assert.Equal(t, expectedOk, ok, "IntersectInts(%v, %v) ok", tt.a, tt.b)
		})
	}
}
