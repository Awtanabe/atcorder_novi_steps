package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(in, &N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	// dp[i]: A[:i] まで考えた総和
	dp := make([]int, N+1)
	// 累積和: dp の和を持っておく
	pref := make([]int, N+1)
	dp[0] = 1 // 空の分割 1 通り
	pref[0] = 1

	// スタックで max と min の寄与を管理
	type pair struct{ val, idx int }

	maxSt, minSt := []pair{}, []pair{}

	for i := 1; i <= N; i++ {
		val := A[i-1]

		// 最大値の寄与
		for len(maxSt) > 0 && maxSt[len(maxSt)-1].val <= val {
			maxSt = maxSt[:len(maxSt)-1]
		}
		l := 0
		if len(maxSt) > 0 {
			l = maxSt[len(maxSt)-1].idx
		}
		// pref[i-1] - pref[l-1] が dp[l..i-1] の総和
		add := (pref[i-1] - pref[l-1] + MOD) % MOD
		dp[i] = (dp[i] + add*val) % MOD
		maxSt = append(maxSt, pair{val, i})

		// 最小値の寄与
		for len(minSt) > 0 && minSt[len(minSt)-1].val >= val {
			minSt = minSt[:len(minSt)-1]
		}
		l = 0
		if len(minSt) > 0 {
			l = minSt[len(minSt)-1].idx
		}
		sub := (pref[i-1] - pref[l-1] + MOD) % MOD
		dp[i] = (dp[i] - sub*val%MOD + MOD) % MOD
		minSt = append(minSt, pair{val, i})

		pref[i] = (pref[i-1] + dp[i]) % MOD
	}

	fmt.Println(dp[N] % MOD)
}
