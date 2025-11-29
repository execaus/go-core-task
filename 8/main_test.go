package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitGroup_DefaultScenario_Success(t *testing.T) {
	goroutineCount := 2
	wg := NewWaitGroup()
	wg.Add(goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			wg.Done()
		}()
	}

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	assert.GreaterOrEqual(t, elapsed.Milliseconds(), int64(50))
}

func TestWaitGroup_AddAfterDone_Panic(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	wg.Done()

	assert.Panics(t, func() { wg.Add(1) })
}

func TestWaitGroup_AddZero_Panic(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(0)

	assert.Panics(t, func() { wg.Done() })
}

func TestWaitGroup_MultipleWaits_GoroutineCorrectUnlocked(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)

	c1Done := make(chan struct{})
	c2Done := make(chan struct{})

	go func() {
		defer close(c1Done)
		wg.Wait()
	}()
	go func() {
		defer close(c2Done)
		wg.Wait()
	}()

	time.Sleep(50 * time.Millisecond)
	wg.Done()

	assert.Eventually(t, func() bool {
		select {
		case <-c1Done:
			return true
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond)

	assert.Eventually(t, func() bool {
		select {
		case <-c2Done:
			return true
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond)
}

func TestWaitGroup_Concurrent_WaitBlocked(t *testing.T) {
	const goroutineCount = 100
	wg := NewWaitGroup()
	wg.Add(goroutineCount)

	finished := make(chan struct{}, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			wg.Done()
			finished <- struct{}{}
		}()
	}

	waitDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitDone)
	}()

	for i := 0; i < goroutineCount; i++ {
		<-finished
	}

	assert.Eventually(t, func() bool {
		select {
		case <-waitDone:
			return true
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond, "Wait не завершился после всех Done")

	assert.NotPanics(t, func() { wg.Wait() }, "Wait после завершения всех Done вызвал панику")
}

func TestWaitGroup_AddNegative_Panic(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)

	assert.Panics(t, func() { wg.Add(-2) })
}
