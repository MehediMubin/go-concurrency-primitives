package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	mp map[int]int
	wg sync.WaitGroup
)

func writeMapWithMutex(key, value int) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	mp[key] = value
}

func writeMapWithoutMutex(key, value int) {
	defer wg.Done()

	mp[key] = value
}

func main() {
	var choice int

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Increment with mutex")
		fmt.Println("2. Increment without mutex")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice (0, 1 or 2): ")

		_, err := fmt.Scan(&choice)
		if err != nil || (choice < 0 || choice > 2) {
			fmt.Println("Invalid choice. Please enter 0, 1 or 2.")
			continue
		}

		if choice == 0 {
			fmt.Println("Exiting the program")
			break
		}

		mp = make(map[int]int)
		fmt.Println("Starting map writing operation...")
	
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			if choice == 1 {
				go writeMapWithMutex(i, i * 2)
			} else {
				go writeMapWithoutMutex(i, i * 2)
			}
		}

		wg.Wait()

		if choice == 1 {
			fmt.Println("Final map (with mutex):")
		} else {
			fmt.Println("Final map (without mutex):")
		}

		for i := 1; i <= 15; i++ {
			fmt.Println(i, mp[i])
		}
	}
}