package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	H := make([]int64, N)
	var total int64
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &H[i])
		total += H[i]
	}

	if K >= N {
		fmt.Println(0)
		return
	}

	// 大きい順に並べる
	sort.Slice(H, func(i, j int) bool { return H[i] > H[j] })

	var topK int64
	for i := 0; i < K; i++ {
		topK += H[i]
	}
	fmt.Println(total - topK)
}
