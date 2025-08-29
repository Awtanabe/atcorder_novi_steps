package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var N int64
	fmt.Fscan(in, &N)

	digs := []int64{3, 5, 7}
	var ans int64

	var dfs func(cur int64, mask int)
	dfs = func(cur int64, mask int) {
		for i, d := range digs {
			nxt := cur*10 + d
			if nxt > N {
				continue
			}
			m := mask | (1 << i)
			if m == 0b111 {
				ans++
			}
			dfs(nxt, m)
		}
	}

	dfs(0, 0)
	fmt.Println(ans)
}
