package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, res int

	fmt.Scan(&N, &K)
	data := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	for i := 0; i < K; i++ {
		res += data[i]
	}

	fmt.Print(res)
}
