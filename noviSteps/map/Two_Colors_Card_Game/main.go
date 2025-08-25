package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)

	blue := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&blue[i])
	}
	fmt.Scan(&M)
	red := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&red[i])
	}
	counts := make(map[string]int)

	for _, e := range blue {
		counts[e]++
	}

	for _, r := range red {
		counts[r]--
	}

	keys := make([]string, len(counts))

	for k := range counts {
		keys = append(keys, k)
	}
	maxval := 0

	for _, v := range counts {
		if maxval < v {
			maxval = v
		}
	}

	fmt.Print(maxval)
}
