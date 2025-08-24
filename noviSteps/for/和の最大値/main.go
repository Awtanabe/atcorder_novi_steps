// https://atcoder.jp/contests/chokudai_S002/tasks/chokudai_S002_c

package main

import (
	"fmt"
)

type Pair struct {
	a int
	b int
}

func (p Pair) Sum() int {
	return p.a + p.b
}

func main() {
	var N, a, b, max int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a, &b)
		p := Pair{a: a, b: b}
		if max < p.Sum() {
			max = p.Sum()
		}
	}
	fmt.Print(max)
}

// bad

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var N, a, b, max int
// 	fmt.Scan(&N)

// 	data := make([][]int, N)

// 	for i, _ := range data {
// 		fmt.Scan(&a, &b)
// 		data[i] = []int{a, b}
// 	}

// 	for _, d := range data {
// 		if max < d[0]+d[1] {
// 			max = d[0] + d[1]
// 		}
// 	}
// 	fmt.Print(max)
// }
