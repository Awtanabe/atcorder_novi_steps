package main

import (
	"fmt"
	"strings"
)

func maxScore(s string) int {
	runes := []rune(s)
	max := 0
	for i := 1; i < len(s); i++ {
		sum := 0
		sum += strings.Count(string(runes[0:i]), "0") + strings.Count(string(runes[i:]), "1")
		if max < sum {
			max = sum
		}
	}

	return max
}

func main() {
	fmt.Println(maxScore("011101"))
}

// 累積和で解く
// 左の数と右の数をintスライスに入れて計算してる

// package main

// import "fmt"

// func maxScore(s string) int {
// 	n := len(s)
// 	leftZeros := make([]int, n+1)
// 	rightOnes := make([]int, n+1)

// 	// 左から累積和（0の数）
// 	for i := 0; i < n; i++ {
// 		leftZeros[i+1] = leftZeros[i]
// 		if s[i] == '0' {
// 			leftZeros[i+1]++
// 		}
// 	}

// 	// 右から累積和（1の数）
// 	for i := n - 1; i >= 0; i-- {
// 		rightOnes[i] = rightOnes[i+1]
// 		if s[i] == '1' {
// 			rightOnes[i]++
// 		}
// 	}

// 	// 各分割位置でスコアを計算
// 	max := 0
// 	for i := 1; i < n; i++ { // 左と右に分かれるように 1〜n-1
// 		score := leftZeros[i] + rightOnes[i]
// 		if score > max {
// 			max = score
// 		}
// 	}

// 	return max
// }

// func main() {
// 	fmt.Println(maxScore("011101")) // 5
// }
