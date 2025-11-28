package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceStrings(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "ordinary case",
			slice1: []string{"a", "b", "c", "d"},
			slice2: []string{"d", "b"},
			want:   []string{"a", "c"},
		},
		{
			name:   "slice1 empty",
			slice1: []string{},
			slice2: []string{"a", "b"},
			want:   []string{},
		},
		{
			name:   "slice2 empty",
			slice1: []string{"a", "b"},
			slice2: []string{},
			want:   []string{"a", "b"},
		},
		{
			name:   "no intersection",
			slice1: []string{"a", "b"},
			slice2: []string{"c", "d"},
			want:   []string{"a", "b"},
		},
		{
			name:   "full overlap",
			slice1: []string{"a", "b"},
			slice2: []string{"a", "b"},
			want:   []string{},
		},
		{
			name:   "unordered slices with overlap",
			slice1: []string{"c", "a", "e", "b"},
			slice2: []string{"a", "d"},
			want:   []string{"c", "e", "b"},
		},
		{
			name:   "unordered slices with full overlap",
			slice1: []string{"z", "y", "x"},
			slice2: []string{"x", "y", "z"},
			want:   []string{},
		},
		{
			name:   "unordered slice2 with partial overlap",
			slice1: []string{"f", "b", "c"},
			slice2: []string{"c", "q"},
			want:   []string{"f", "b"},
		},
		{
			name:   "unordered slice1 and slice2, no overlap",
			slice1: []string{"o", "t", "h"},
			slice2: []string{"f", "i"},
			want:   []string{"o", "t", "h"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceStrings(tt.slice1, tt.slice2)
			assert.Equal(t, tt.want, got, "DifferenceStrings()")
		})
	}
}
