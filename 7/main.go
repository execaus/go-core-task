package main

import (
	"sync"
)

type WaitGroup struct {
	m        sync.Mutex
	count    int
	done     bool
	notifier chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		count:    0,
		notifier: make(chan struct{}),
	}
}

// Add увеличивает счетчик WaitGroup на значение value.
func (g *WaitGroup) Add(value int) {
	g.m.Lock()
	defer g.m.Unlock()

	if g.done {
		panic("wait group done")
	}

	if g.count+value < 0 {
		panic("invalid value")
	}

	g.count += value
}

// Done уменьшает счетчик WaitGroup на 1.
func (g *WaitGroup) Done() {
	g.m.Lock()
	defer g.m.Unlock()

	if g.count == 0 {
		panic("count is zero")
	}

	g.count -= 1

	if g.count == 0 {
		g.done = true
		close(g.notifier)
	}
}

// Wait блокирует выполнение до тех пор, пока счетчик count не станет равен 0.
func (g *WaitGroup) Wait() {
	<-g.notifier
}
