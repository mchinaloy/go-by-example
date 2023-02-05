package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := GetAverage(numbers)
	fmt.Println(result)
}

func GetAverage(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return 0
	}
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total / length
}
