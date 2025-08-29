package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int64 = 998244353

type node struct {
	val, w int64 // 区間最大(最小)が val になる j の重み合計 w = sum dp[j-1]
}

func add(x, y int64) int64 {
	x += y
	if x >= MOD {
		x -= MOD
	}
	return x
}
func sub(x, y int64) int64 {
	x -= y
	if x < 0 {
		x += MOD
	}
	return x
}
func mul(x, y int64) int64 {
	return (x % MOD) * (y % MOD) % MOD
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(in, &N)

	A := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	// dpPrev = dp[i-1]
	dpPrev := int64(1)

	var cMax, cMin int64     // それぞれ Σ dp[j-1]*rangeMax(j..i), Σ dp[j-1]*rangeMin(j..i)
	stMax := make([]node, 0) // 値は単調↓
	stMin := make([]node, 0) // 値は単調↑

	var dpi int64
	for i := 0; i < N; i++ {
		w := dpPrev

		// --- 最大値側（単調減少スタック）
		weight := w
		for len(stMax) > 0 && stMax[len(stMax)-1].val <= A[i] {
			top := stMax[len(stMax)-1]
			stMax = stMax[:len(stMax)-1]
			cMax = sub(cMax, mul(top.val, top.w))
			weight = add(weight, top.w)
		}
		stMax = append(stMax, node{A[i], weight})
		cMax = add(cMax, mul(A[i], weight))

		// --- 最小値側（単調増加スタック）
		weight = w
		for len(stMin) > 0 && stMin[len(stMin)-1].val >= A[i] {
			top := stMin[len(stMin)-1]
			stMin = stMin[:len(stMin)-1]
			cMin = sub(cMin, mul(top.val, top.w))
			weight = add(weight, top.w)
		}
		stMin = append(stMin, node{A[i], weight})
		cMin = add(cMin, mul(A[i], weight))

		dpi = sub(cMax, cMin) // j=i の項は (A[i]-A[i])=0 なので自然に相殺される
		dpPrev = dpi
	}

	fmt.Println((dpi%MOD + MOD) % MOD)
}
