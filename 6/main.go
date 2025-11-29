package main

import (
	"math/rand"
)

type UnbufferRandomGenerator struct {
	maxRandomValue int
	notifier       chan int
}

// NewUnbufferRandomGenerator создаёт новый генератор случайных чисел
// с неблокирующим (небуферизированным) каналом для отправки значений.
func NewUnbufferRandomGenerator(maxRandomValue int) *UnbufferRandomGenerator {
	return &UnbufferRandomGenerator{
		maxRandomValue: maxRandomValue,
		notifier:       make(chan int),
	}
}

// Channel возвращает канал, по которому генератор отправляет значения.
func (g *UnbufferRandomGenerator) Channel() chan int {
	return g.notifier
}

// Generate генерирует случайное число в диапазоне [0, maxRandomValue)
// и отправляет его в канал. Вызов блокируется, пока получатель не прочитает значение.
func (g *UnbufferRandomGenerator) Generate() {
	g.notifier <- rand.Intn(g.maxRandomValue)
}
