package main

type UnbufferRandomGenerator struct {
	notifier       chan int
	maxRandomValue int
}

func NewUnbufferRandomGenerator(maxRandomValue int) *UnbufferRandomGenerator {
	return &UnbufferRandomGenerator{}
}

func (g *UnbufferRandomGenerator) Channel() chan int {
	// TODO
}

func (g *UnbufferRandomGenerator) Generate() {
	// TODO
}
