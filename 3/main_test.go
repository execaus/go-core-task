package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_Valid_Success(t *testing.T) {
	m := NewStringIntMap()

	m.Add("one", 1)
	m.Add("two", 2)

	val, ok := m.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestGet_Valid_Success(t *testing.T) {
	m := NewStringIntMap()

	m.Add("one", 1)
	m.Add("two", 2)

	val, ok := m.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = m.Get("two")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
}

func TestExists_Exists_Success(t *testing.T) {
	m := NewStringIntMap()

	m.Add("one", 1)

	assert.True(t, m.Exists("one"))
	assert.False(t, m.Exists("two"))
}

func TestRemove_Valid_Success(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	m.Remove("one")

	assert.False(t, m.Exists("one"))
}

func TestCopy_Valid_Success(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	copied := m.Copy()

	assert.Equal(t, 2, len(copied))
	assert.Equal(t, 1, copied["one"])
	assert.Equal(t, 2, copied["two"])
}

func TestGet_NonExistentKey_Success(t *testing.T) {
	m := NewStringIntMap()

	val, ok := m.Get("missing")

	assert.False(t, ok)
	assert.Equal(t, 0, val)
}
