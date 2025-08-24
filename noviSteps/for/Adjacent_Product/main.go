package main

import (
	"fmt"
)

func main() {
	var N, val int

	fmt.Scan(&N)

	data := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&val)
		data[i] = val
	}

	for i, _ := range data {
		if len(data) > i+1 {
			fmt.Printf("%d ", data[i]*data[i+1])
		}
	}
}

