package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		increment(&numbers[0])
	}
	fmt.Println(numbers[0])
	wg.Wait()
}

func increment(number *int) {
	mu.Lock()
	defer mu.Unlock()
	*number++
	wg.Done()
}
