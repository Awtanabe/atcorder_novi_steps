package main

import (
	"fmt"
)

func main() {
	var N, tmp int
	fmt.Scan(&N)

	// 0 ~ 2000なので 2001個必要
	// []boolでも良いらしい！
	bucket := make([]int, 2001)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		bucket[tmp] = 1
	}

	for i, d := range bucket {
		if d != 1 {
			fmt.Print(i)
			break
		}
	}

}
