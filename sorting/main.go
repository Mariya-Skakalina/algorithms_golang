package main

import "fmt"

func selectionSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}

		(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
	}
}

func main() {
	arr := []int{12, 5, 2, 98, 75, 0, 3, 1}
	selectionSort(&arr)
	fmt.Println(arr)
}
