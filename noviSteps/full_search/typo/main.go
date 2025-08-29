package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	if S == T {
		fmt.Println("Yes")
		return
	}

	n := len(S)
	diffs := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if S[i] != T[i] {
			diffs = append(diffs, i)
			if len(diffs) > 2 {
				fmt.Println("No")
				return
			}
		}
	}

	// 異なる箇所が2つ、かつ隣接、かつ入れ替えで一致するか
	if len(diffs) == 2 && diffs[1] == diffs[0]+1 {
		i, j := diffs[0], diffs[1]
		if S[i] == T[j] && S[j] == T[i] {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
