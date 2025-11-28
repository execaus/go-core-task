package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomSlice_NegativeSize_Error(t *testing.T) {
	const size = -1

	s, err := randomSlice(size)

	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestRandomSlice_ZeroSize_Error(t *testing.T) {
	const size = 0

	s, err := randomSlice(size)

	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestRandomSlice_ManyCalls_DifferentValues(t *testing.T) {
	const size = 10

	first, err := randomSlice(size)
	assert.NoError(t, err)

	second, err := randomSlice(size)
	assert.NoError(t, err)

	assert.Equal(t, size, len(first))
	assert.Equal(t, size, len(second))
	assert.NotEqual(t, first, second)
}

func TestSliceExample_EmptySlice_Error(t *testing.T) {
	var in []int

	result, err := sliceExample(in)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSliceExample_CorrectSlice_OnlyEvenNumbers(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result, err := sliceExample(in)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestAddElements_CorrectInput_Success(t *testing.T) {
	in := []int{2, 3, 4}
	num := 5
	expected := []int{2, 3, 4, 5}

	result, err := addElements(in, num)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestAddElements_SliceNil_Error(t *testing.T) {
	var in []int
	num := 5

	result, err := addElements(in, num)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCopySlice_SliceNil_Error(t *testing.T) {
	var in []int

	result, err := copySlice(in)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCopySlice_CorrectSlice_Success(t *testing.T) {
	in := []int{1, 2, 3}

	result, err := copySlice(in)

	assert.NoError(t, err)
	assert.Equal(t, in, result)
}

func TestCopySlice_ModificationsToOriginal_DoNotAffectCopy(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	result, err := copySlice(in)
	in[0] = 100

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestRemoveElement_SliceNil_Error(t *testing.T) {
	var in []int
	index := 0

	result, err := removeElement(in, index)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestRemoveElement_NegativeIndex_Error(t *testing.T) {
	in := []int{1, 2, 3}
	index := -1

	result, err := removeElement(in, index)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestRemoveElement_OutOfRangeIndex_Error(t *testing.T) {
	in := []int{1, 2, 3}
	index := 3

	result, err := removeElement(in, index)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestRemoveElement_CorrectInput_Success(t *testing.T) {
	in := []int{10, 20, 30, 40, 50}
	index := 2
	expected := []int{10, 20, 40, 50}

	result, err := removeElement(in, index)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
