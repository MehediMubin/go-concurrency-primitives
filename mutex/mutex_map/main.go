package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	mp = make(map[int]int)
	wg sync.WaitGroup
)

func writeMapWithMutex(key, value int) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	mp[key] = value
}

func main() {
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go writeMapWithMutex(i, i * 2)	
	}

	wg.Wait()

	for i := 1; i <= 10; i++ {
		fmt.Println(i, mp[i])
	}
}