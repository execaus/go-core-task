package main

import (
	"log"
	"math"
	"math/rand"
)

const (
	numberCount = 16
)

type Result struct {
	In  uint8
	Out float64
}

func main() {
	pipeline := NewPipeline()
	numbers := randomNumbers(numberCount)

	go read(pipeline)

	for _, number := range numbers {
		pipeline.In() <- number
	}

	pipeline.Close()
}

func randomNumbers(count int) []uint8 {
	numbers := make([]uint8, count)

	for i := 0; i < count; i++ {
		numbers[i] = uint8(rand.Intn(math.MaxUint8))
	}

	return numbers
}

func read(pipeline *Pipeline) {
	for result := range pipeline.Out() {
		log.Printf("%d => %f\n", result.In, result.Out)
	}
}

type Pipeline struct {
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) In() chan<- uint8 {
	// TODO
}

func (p *Pipeline) Out() <-chan *Result {
	// TODO
}

func (p *Pipeline) Close() {
	// TODO
}
