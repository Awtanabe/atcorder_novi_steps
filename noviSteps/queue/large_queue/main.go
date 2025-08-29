package main

import (
	"bufio"
	"fmt"
	"os"
)

type run struct {
	x int64 // 値
	c int64 // 個数
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &Q)

	queue := make([]run, 0, Q)
	head := 0

	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var c, x int64
			fmt.Fscan(in, &c, &x)
			queue = append(queue, run{x: x, c: c})
		} else { // t == 2
			var k int64
			fmt.Fscan(in, &k)
			var sum int64
			for k > 0 {
				r := &queue[head]
				if r.c <= k {
					sum += r.x * r.c
					k -= r.c
					head++
				} else {
					sum += r.x * k
					r.c -= k
					k = 0
				}
			}
			fmt.Fprintln(out, sum)
			// メモリ節約のため、たまにスライスを前詰めしても良いが必須ではない
			if head > 1<<20 && head*2 > len(queue) {
				queue = append([]run(nil), queue[head:]...)
				head = 0
			}
		}
	}
}
