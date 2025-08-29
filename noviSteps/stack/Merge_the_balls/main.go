package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	stack := make([]int64, 0, N)
	for i := 0; i < N; i++ {
		var a int64
		fmt.Fscan(in, &a)
		// 新しいボール（指数a）を右端に追加
		stack = append(stack, a)
		// 右端から2つが同じ指数なら統合して指数+1を再度右端に追加
		for len(stack) >= 2 {
			n := len(stack)
			if stack[n-1] != stack[n-2] {
				break
			}
			x := stack[n-1] + 1
			stack = stack[:n-2]      // 2つ取り除く
			stack = append(stack, x) // 和のボール（指数+1）を追加
		}
	}
	fmt.Println(len(stack))
}
