package main

import "testing"

func Test_GetAverage_Returns_Average(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := GetAverage(numbers)
	if result != 3 {
		t.Error("Expected", 3, "Got", result)
	}
}

func Test_GetAverage_WithNoNumbers_Returns_Zero(t *testing.T) {
	numbers := []int{}
	result := GetAverage(numbers)
	if result != 0 {
		t.Error("Expected", 0, "Got", result)
	}
}
