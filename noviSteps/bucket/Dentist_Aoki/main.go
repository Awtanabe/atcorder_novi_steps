package main

import "fmt"

func main() {
	var N, Q, t, c int
	fmt.Scan(&N, &Q)

	teeth := make([]bool, N)
	for i := range teeth {
		teeth[i] = true
	}

	for i := 0; i < Q; i++ {
		fmt.Scan(&t)
		// tは 1~ Nなので out of rangeする
		idx := t - 1
		if teeth[idx] {
			teeth[idx] = false
		} else {
			teeth[idx] = true
		}
	}
	for _, d := range teeth {
		if d == true {
			c++
		}
	}
	fmt.Print(c)
}
