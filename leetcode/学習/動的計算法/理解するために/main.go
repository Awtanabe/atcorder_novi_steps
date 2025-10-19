package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	// 5 までのインデックス
	// 0 1 2 3 4 5
	// 1 1 2 3 5 8
	fmt.Println(fib(5))
}
