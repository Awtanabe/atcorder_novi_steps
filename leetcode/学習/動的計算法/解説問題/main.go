package main

// 再帰のフィボナッチ

// package main

// import "fmt"

// func fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}

// 	return fib(n-2) + fib(n-1)
// }

// func main() {
// 	// 5 までのインデックス
// 	// 0 1 2 3 4 5
// 	// 1 1 2 3 5 8
// 	fmt.Println(fib(5))
// }

// 貰う すでに記録してる結果をもらうから
import "fmt"

func fib(n int) int {
	// n+1個の要素を持つスライスを作成（初期値はすべて0）
	memo := make([]int, n+1)

	// 0番目と1番目の初期値を設定
	memo[0] = 1
	memo[1] = 1

	// 2番目以降のフィボナッチ数を順に計算
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
}

// 配る
// package main

// import "fmt"

// func fib(N int) int {
// 	// memo[i]: i番目のフィボナッチ数
// 	memo := make([]int, N+1)

// 	// 0番目を初期化（1番目はループ中で決まる）
// 	memo[0] = 1

// 	for i := 0; i < N; i++ {
// 		// i番目の値を i+1 番目と i+2 番目に足す（伝播させる）
// 		memo[i+1] += memo[i]
// 		if i+2 < N+1 {
// 			memo[i+2] += memo[i]
// 		}
// 	}

// 	return memo[N]
// }

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Printf("fib(%d) = %d\n", i, fib(i))
// 	}
// }
