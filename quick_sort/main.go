package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	arr = append(arr[:pivotIndex], arr[pivotIndex+1:]...)

	left := make([]int, 0)
	right := make([]int, 0)

	for _, value := range arr {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	array := []int{3, 6, 8, 10, 1, 2, 1}
	sortedArray := quickSort(array)
	fmt.Println(sortedArray)
}
