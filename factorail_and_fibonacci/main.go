package main

import "fmt"

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	x = x * factorial(x-1)
	return x
}

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	x = fibonacci(x-1) + fibonacci(x-2)
	return x
}

func main() {
	x := 4
	fmt.Printf("Факториал числа %d, равен %d\n", x, factorial(x))
	fmt.Printf("Число Фибоначчи под номером %d равно %d\n", x, fibonacci(x))
}
