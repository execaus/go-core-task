package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipeline_CorrectScenario_CubeConversion(t *testing.T) {
	p := NewPipeline()

	inputs := []uint8{0, 1, 2, 3, 4, 5, 10, 255}
	out := p.Out()

	go func() {
		for _, n := range inputs {
			p.In() <- n
		}
		p.Close()
	}()

	results := collectResults(out)
	require.Len(t, results, len(inputs))

	for i, r := range results {
		expected := float64(r.In) * float64(r.In) * float64(r.In)
		assert.Equal(t, inputs[i], r.In)
		assert.Equal(t, expected, r.Out)
	}
}

func TestPipeline_CorrectScenario_ChannelClosed(t *testing.T) {
	p := NewPipeline()
	p.Close()

	_, ok := <-p.Out()
	assert.False(t, ok)
}

func TestPipeline_CorrectScenario_CorrectOrder(t *testing.T) {
	p := NewPipeline()

	inputs := []uint8{2, 4, 6, 8}
	out := p.Out()

	go func() {
		for _, n := range inputs {
			p.In() <- n
		}
		p.Close()
	}()

	results := collectResults(out)

	for i, r := range results {
		assert.Equal(t, inputs[i], r.In)
		expected := float64(inputs[i]) * float64(inputs[i]) * float64(inputs[i])
		assert.Equal(t, expected, r.Out)
	}
}

func collectResults(ch <-chan *Result) []*Result {
	var results []*Result
	for r := range ch {
		results = append(results, r)
	}
	return results
}
