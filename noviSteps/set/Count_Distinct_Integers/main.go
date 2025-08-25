package main

import (
	"fmt"
)

func main() {
	var N, key int
	set := make(map[int]struct{})

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&key)
		set[key] = struct{}{}
	}

	fmt.Print(len(set))
}
