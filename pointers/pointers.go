package main

import "fmt"

func main() {
	number := 10
	fmt.Println(number)
	increment(&number)
	fmt.Println(number)
}

func increment(number *int) {
	*number++
	fmt.Println(*number)
}
