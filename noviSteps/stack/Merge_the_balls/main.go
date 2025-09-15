// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var N int
// 	fmt.Fscan(in, &N)

// 	stack := make([]int64, 0, N)
// 	for i := 0; i < N; i++ {
// 		var a int64
// 		fmt.Fscan(in, &a)
// 		// 新しいボール（指数a）を右端に追加
// 		stack = append(stack, a)
// 		// 右端から2つが同じ指数なら統合して指数+1を再度右端に追加
// 		for len(stack) >= 2 {
// 			n := len(stack)
// 			if stack[n-1] != stack[n-2] {
// 				break
// 			}
// 			x := stack[n-1] + 1
// 			stack = stack[:n-2]      // 2つ取り除く
// 			stack = append(stack, x) // 和のボール（指数+1）を追加
// 		}
// 	}
// 	fmt.Println(len(stack))
// }

package main

// 条件 列にあるボールが1つなら操作を終了 len(data) == 0
// 右から1, 2番目のものの大きさが異なるなら操作終了
// 右から一番目と2番目の大きさが等しいなら 2つのボールを取り除き && それらの合計の入れる

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	data := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	res := make([]int, 0, N)

	for i := 0; i < N; i++ {
		res = append(res, data[i])

		if len(res) < 2 {
			continue
		}

		for len(res) >= 2 {
			if res[len(res)-1] != res[len(res)-2] {
				break
			}
			tmp := res[len(res)-2]
			res = res[:len(res)-2]
			res = append(res, tmp+1)
		}
	}

	fmt.Println(len(res))
	fmt.Println(res)
}
