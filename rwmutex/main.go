// The following code was written by - https://github.com/zahansafallwa1511

package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct{
	mu sync.RWMutex
	value int
}

func (c *Counter) Increment(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
	fmt.Printf("Writer %d incremented the counter to: %d\n", id, c.value)
}

func (c *Counter) Value(id int) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	fmt.Printf("Reader %d is reading the counter: %d\n", id, c.value)
	time.Sleep(500 * time.Millisecond) // Simulate a longer read delay
	fmt.Printf("Reader %d finished reading\n", id)
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	readerCount := 0
	writerCount := 0

	for {
		fmt.Println("\nEnter 1 to spawn multiple readers, 2 to increment, 0 to exit:")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			for range 5 {
				readerCount++
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					counter.Value(id)
				}(readerCount)
			}
		case 2:
			writerCount++
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				counter.Increment(id)
			}(writerCount)
		case 0:
			wg.Wait()
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Invalid choice, please enter 1, 2, or 0.")
		}
	}
}