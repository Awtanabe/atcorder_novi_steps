package main

import "fmt"

func rec(n int) {
	if n == 0 {
		fmt.Println("終了")
		return
	}

	fmt.Printf("前処理: n=%d\n", n)
	rec(n - 1) // スタックフレームに積み上げていく
	// 関数が優先で呼び出した関数が終了(再帰の場合はreturnで終わらせる)したら、呼び出しの直後から再開
	// 呼ばれた関数が終わるまで待機する。 通常の関数でもやっているけど別関数だから違和感がないのか
	fmt.Printf("後処理: n=%d\n", n)
}

func main() {
	rec(2)
}
