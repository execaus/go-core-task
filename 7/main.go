package main

import "sync"

// Merge принимает несколько каналов с int и объединяет их в один канал.
func Merge(channels ...<-chan int) <-chan int {
	if len(channels) == 0 {
		panic("channels not found")
	}

	result := make(chan int)

	wgClosed := sync.WaitGroup{}

	wgClosed.Add(len(channels))

	for _, channel := range channels {
		go func(ch <-chan int) {
			for v := range ch {
				result <- v
			}
			wgClosed.Done()
		}(channel)
	}

	go func() {
		wgClosed.Wait()
		close(result)
	}()

	return result
}
