package main

import (
	"fmt"
)

func sum(n []int) int {

	if len(n) == 0 {
		return 0
	}

	// この行が感覚的にまだ理解できない
	return n[0] + sum(n[1:])
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
}
