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

	for x := N; x <= 919; x++ { // 制約上ここまで見れば必ず見つかる
		a := x / 100       // 百の位
		b := (x / 10) % 10 // 十の位
		c := x % 10        // 一の位
		if a*b == c {
			fmt.Println(x)
			return
		}
	}
}
