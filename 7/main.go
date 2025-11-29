package main

type WaitGroup struct{}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{}
}

func (g *WaitGroup) Add(value int) {

}

func (g *WaitGroup) Done() {
}

func (g *WaitGroup) Wait() {
}
