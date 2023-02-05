package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{}
	if resp, err := check_numbers(numbers); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func check_numbers(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("No numbers!")
	}
	return len(numbers), nil
}
