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
	in  chan uint8
	out chan *Result
}

func NewPipeline() *Pipeline {
	p := Pipeline{
		in:  make(chan uint8),
		out: make(chan *Result),
	}
	go p.process()
	return &p
}

// In возвращает канал для отправки входных чисел типа uint8 в pipeline.
func (p *Pipeline) In() chan<- uint8 {
	return p.in
}

// Out возвращает канал для получения результатов обработки чисел.
func (p *Pipeline) Out() <-chan *Result {
	return p.out
}

// Close закрывает входной канал pipeline.
func (p *Pipeline) Close() {
	close(p.in)
}

func (p *Pipeline) process() {
	for value := range p.in {
		p.out <- &Result{
			In:  value,
			Out: math.Pow(float64(value), 3),
		}
	}
	close(p.out)
}
