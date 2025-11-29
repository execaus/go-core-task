package main

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUnbufferRandomGenerator_CorrectMaxValue_CorrectFields(t *testing.T) {
	maxValue := 100

	g := NewUnbufferRandomGenerator(maxValue)

	require.NotNil(t, g)
	assert.Equal(t, maxValue, g.maxRandomValue)
	require.NotNil(t, g.notifier)
}

func TestChannel_CorrectCreate_Success(t *testing.T) {
	g := NewUnbufferRandomGenerator(math.MaxInt)

	ch := g.Channel()

	require.NotNil(t, ch)
	assert.Equal(t, g.notifier, ch)
}

func TestGenerate_GenerateNumber_Success(t *testing.T) {
	g := NewUnbufferRandomGenerator(math.MaxInt)

	go g.Generate()

	select {
	case v := <-g.Channel():
		assert.GreaterOrEqual(t, v, 0)
		assert.Less(t, v, g.maxRandomValue)
	case <-time.After(time.Second):
		require.Fail(t, "timeout: no value received from notifier channel")
	}
}

func TestGenerate_TwoGenerate_RandomNumbers(t *testing.T) {
	var result [2]int
	g := NewUnbufferRandomGenerator(math.MaxInt)

	go g.Generate()
	go g.Generate()

	for i := 0; i < 2; i++ {
		select {
		case result[i] = <-g.Channel():
		case <-time.After(time.Second):
			require.Fail(t, "timeout: no value received from notifier channel")
		}
	}

	assert.NotEqual(t, result[0], result[1])
}
