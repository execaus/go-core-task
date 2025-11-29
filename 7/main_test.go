package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMerge_TwoChannels_Success(t *testing.T) {
	c1 := makeChannels(1, 2, 3)
	c2 := makeChannels(4, 5, 6)

	merged := Merge(c1, c2)

	var results []int
	for v := range merged {
		results = append(results, v)
	}

	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, results)
}

func TestMerge_OneChannel_Success(t *testing.T) {
	c := makeChannels(10, 20, 30)

	merged := Merge(c)

	var results []int
	for v := range merged {
		results = append(results, v)
	}

	assert.Equal(t, []int{10, 20, 30}, results)
}

func TestMerge_NoChannels_Panic(t *testing.T) {
	assert.Panics(t, func() { Merge() })
}

func TestMerge_EmptyChannel_Success(t *testing.T) {
	c1 := makeChannels()
	c2 := makeChannels(7, 8)

	merged := Merge(c1, c2)

	results := []int{}
	for v := range merged {
		results = append(results, v)
	}

	assert.ElementsMatch(t, []int{7, 8}, results)
}

func TestMerge_SlowChannel_Success(t *testing.T) {
	c1 := make(chan int)
	c2 := makeChannels(1, 2)

	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- 99
		close(c1)
	}()

	merged := Merge(c1, c2)

	var results []int
	for v := range merged {
		results = append(results, v)
	}

	assert.ElementsMatch(t, []int{1, 2, 99}, results)
}

func makeChannels(values ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range values {
			c <- v
		}
		close(c)
	}()
	return c
}
