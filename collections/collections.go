package main

import "fmt"

func main() {
	create_arrays()
	create_slices()
	create_maps()
}

func create_arrays() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	numbers_using_make := make([]int, 5)
	numbers_using_make[0] = 21
	fmt.Println(numbers_using_make)
}

func create_slices() {
	numbers := []int{}
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Println(numbers)
	fmt.Println(numbers[:2])

	numbers = append(numbers, 13)
	fmt.Println(len(numbers))
	remove_from_slice(numbers, 1)
	fmt.Println(numbers)
}

func remove_from_slice(numbers []int, index int) []int {
	return append(numbers[:index], numbers[index+1:]...)
}

func create_maps() {
	names_numbers := map[string]int{"Michael": 21}
	fmt.Println(names_numbers)

	names_numbers_using_make := make(map[string]int, 1)
	names_numbers_using_make["Michael"] = 21
	fmt.Println(names_numbers_using_make)

	if v, ok := names_numbers["Matthew"]; ok {
		fmt.Println(v)
	}
}
