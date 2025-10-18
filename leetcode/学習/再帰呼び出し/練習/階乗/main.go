package main

import "fmt"

// 階乗は 1から nまでの数字を足し合わせたもの
// 0! = 1 0の階乗は1になる
func kaijyo(n int) int {
	if n == 0 {
		return 1
	}

	// n 5
	// n 4 ...
	// n == 0 は1
	return n * kaijyo(n-1)
}

func main() {

	fmt.Println(kaijyo(5))
}
