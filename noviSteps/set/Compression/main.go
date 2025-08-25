package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	d := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&d[i])
	}

	m := make(map[int]struct{})

	for _, dd := range d {
		m[dd] = struct{}{}
	}

	var sd []int
	for key, _ := range m {
		sd = append(sd, key)
	}

	sort.Slice(sd, func(i, j int) bool {
		return sd[i] < sd[j]
	})

	fmt.Println(len(sd))

	for _, x := range sd {
		fmt.Print(x)
		fmt.Print(" ")
	}
}
