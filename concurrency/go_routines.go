package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	wg.Add(1)
	go print_numbers(numbers)
	wg.Wait()
}

func print_numbers(numbers []int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
	wg.Done()
}
