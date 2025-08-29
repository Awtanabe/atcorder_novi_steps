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

	var sum, sq int64
	for i := 0; i < N; i++ {
		var x int64
		fmt.Fscan(in, &x)
		sum += x
		sq += x * x
	}

	// (sum^2 - sum of squares) / 2
	ans := (sum*sum - sq) / 2
	fmt.Println(ans)
}
