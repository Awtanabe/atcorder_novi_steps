package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	A := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	var Q int
	fmt.Fscan(in, &Q)

	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var k int
			var x int64
			fmt.Fscan(in, &k, &x)
			A[k-1] = x
		} else { // t == 2
			var k int
			fmt.Fscan(in, &k)
			fmt.Fprintln(out, A[k-1])
		}
	}
}
