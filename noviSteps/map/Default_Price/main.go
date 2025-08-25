package main

// https://atcoder.jp/contests/abc308/tasks/abc308_b

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	// 食べた皿の色
	C := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&C[i])
	}

	// 特別価格が設定されている皿の色
	D := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&D[i])
	}

	// 価格 P0, P1..PM
	P := make([]int, M+1)
	for i := 0; i <= M; i++ {
		fmt.Scan(&P[i])
	}

	// 色→価格の対応表を作る
	price := make(map[string]int)
	for i := 0; i < M; i++ {
		price[D[i]] = P[i+1] // D[i] の価格は P[i+1]
	}

	// 合計計算
	total := 0
	for _, color := range C {
		if v, ok := price[color]; ok {
			total += v
		} else {
			total += P[0] // デフォルト価格
		}
	}

	fmt.Println(total)
}
