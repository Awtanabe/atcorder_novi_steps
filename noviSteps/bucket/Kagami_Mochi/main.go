package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// 直径は 1..100 の範囲なので 101 要素のバケットを用意
	var bucket [101]int

	for i := 0; i < N; i++ {
		var d int
		fmt.Scan(&d)
		bucket[d]++ // 直径 d の餅が何枚あるか数える
	}

	// 出現した直径の種類数を数える
	ans := 0
	for d := 1; d <= 100; d++ {
		if bucket[d] > 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
