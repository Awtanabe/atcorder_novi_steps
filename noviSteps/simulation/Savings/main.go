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

	lo, hi := int64(0), int64(1)
	// 上限をざっくり広げる（hi は必ず条件を満たすように）
	for hi*(hi+1)/2 < N {
		hi *= 2
	}

	// lo は未満, hi は条件を満たす最小値に収束
	for lo+1 < hi {
		mid := (lo + hi) / 2
		if mid*(mid+1)/2 >= N {
			hi = mid
		} else {
			lo = mid
		}
	}
	fmt.Println(hi)
}
