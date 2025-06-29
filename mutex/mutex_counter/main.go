package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func incrementWithMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func incrementWithoutMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup
	var choice int

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Increment without mutex")
		fmt.Println("2. Increment with mutex")
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

		counter = 0
		fmt.Println("Starting increment operation...")
	
		for range 10 {
			wg.Add(1)
			if choice == 1 {
				go incrementWithoutMutex(&wg)
			} else {
				go incrementWithMutex(&wg)
			}
		}

		wg.Wait()
		if choice == 1 {
			fmt.Println("Final Counter (without mutex):", counter)
		} else {
				fmt.Println("Final Counter (with mutex):", counter)
		}
	}
}