package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	hight := len(arr) - 1

	for low <= hight {
		mid := low + (hight-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			hight = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 3
	result := binarySearch(arr, target)
	if result == -1 {
		fmt.Println("Элемент не найден в массиве")
	} else {
		fmt.Printf("Элемент найден в индексе %d\n", result)
	}

}
