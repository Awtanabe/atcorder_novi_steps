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

	ans := int(1 << 60) // 十分大きい値
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		c := 0
		for x%2 == 0 {
			x /= 2
			c++
		}
		if c < ans {
			ans = c
		}
	}
	fmt.Println(ans)
}
