package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatVariablesWithTypes_ValidInput_Success(t *testing.T) {
	input := []any{123, 83, 255, 3.14, "hello", true, complex(1, 2)}
	expected := []string{
		"Value: 123, Type: int",
		"Value: 83, Type: int",
		"Value: 255, Type: int",
		"Value: 3.14, Type: float64",
		"Value: hello, Type: string",
		"Value: true, Type: bool",
		"Value: (1+2i), Type: complex128",
	}

	result, err := formatVariablesWithTypes(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestFormatVariablesWithTypes_EmptyInput_Error(t *testing.T) {
	input := []any{}

	result, err := formatVariablesWithTypes(input)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCombineVariables_CorrectInput_Success(t *testing.T) {
	input := []any{123, 83, 255, 3.14, "hello", true, complex(1, 2)}
	expected := "123 83 255 3.14 hello true (1+2i)"

	result, err := combineVariables(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCombineVariables_EmptyInput_Error(t *testing.T) {
	input := []any{}
	expected := ""

	result, err := combineVariables(input)

	assert.Error(t, err)
	assert.Equal(t, expected, result)
}

func TestHashVariablesWithSalt_ConsistentHash_SameInputSameHash(t *testing.T) {
	vars := []rune("123 83 255 3.14 hello true (1+2i)")
	salt := "mysalt"

	hash1, err1 := hashVariablesWithSalt(vars, salt)
	hash2, err2 := hashVariablesWithSalt(vars, salt)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, hash1, hash2)
}

func TestHashVariablesWithSalt_DifferentSalt_DifferentHash(t *testing.T) {
	vars := []rune("123 83 255 3.14 hello true (1+2i)")
	salt1 := "salt1"
	salt2 := "salt2"

	hash1, err1 := hashVariablesWithSalt(vars, salt1)
	hash2, err2 := hashVariablesWithSalt(vars, salt2)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEqual(t, hash1, hash2)
}

func TestHashVariablesWithSalt_EmptyVars_Error(t *testing.T) {
	vars := []rune{}
	salt := "salt"

	hash, err := hashVariablesWithSalt(vars, salt)

	assert.Error(t, err)
	assert.Empty(t, hash)
}

func TestHashVariablesWithSalt_EmptySalt_Error(t *testing.T) {
	vars := []rune("123 83")
	salt := ""

	hash, err := hashVariablesWithSalt(vars, salt)

	assert.Error(t, err)
	assert.Empty(t, hash)
}
