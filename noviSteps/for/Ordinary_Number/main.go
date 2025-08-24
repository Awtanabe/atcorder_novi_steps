package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, c int

	fmt.Scan(&N)

	data := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i+2 < N; i++ {
		tmp := []int{data[i], data[i+1], data[i+2]}

		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] > tmp[j]
		})

		if tmp[1] == data[i+1] {
			c++
		}
	}
	fmt.Print(c)
}
