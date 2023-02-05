package main

import "fmt"

func main() {
	var number int = 10
	fmt.Println(number)
	number = increment(number)
	fmt.Println(number)

	numbers := []int{1, 2, 3, 4}
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)

	print_numbers(numbers)

	func() {
		number--
	}()

	fmt.Println(number)

	var decrement = func() {
		number--
	}
	decrement()

	executeFunction(number, func(int) int {
		number = number + 1
		return number
	})
}

func executeFunction(number int, a_func func(int) int) int {
	return a_func(number)
}

func increment(number int) int {
	number = number + 1
	return number
}

func print_numbers(numbers []int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}
