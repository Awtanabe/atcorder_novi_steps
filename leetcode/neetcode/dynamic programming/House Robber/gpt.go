package main

import "fmt"

// dfs(0)
// │
// ├─ dfs(1)
// │   │
// │   ├─ dfs(2)
// │   │   │
// │   │   ├─ dfs(3)
// │   │   │   │
// │   │   │   ├─ dfs(4)
// │   │   │   │   ├─ dfs(5) → 0
// │   │   │   │   └─ dfs(6) → 0
// │   │   │   │        ↳ dfs(4)=1  ✅ memo[4]に保存
// │   │   │   └─ dfs(5) → 0
// │   │   │        ↳ dfs(3)=3  ✅ memo[3]に保存
// │   │   └─ dfs(4)=1  ⚡ (memoから取得)
// │   │        ↳ dfs(2)=10 ✅ memo[2]に保存
// │   └─ dfs(3)=3  ⚡ (memoから取得)
// │        ↳ dfs(1)=10 ✅ memo[1]に保存
// └─ dfs(2)=10 ⚡ (memoから取得)
//      ↳ dfs(0)=12 ✅ memo[0]に保存

// dfs(5) → 範囲外 → 0
// dfs(6) → 範囲外 → 0
// nums[4] → 1

// ステップ1：  dfs(0)
//              /     \
//          dfs(1)     dfs(2)
//          /   \        /  \
//      dfs(2) dfs(3) dfs(3) dfs(4)
//       ↓        ↓      ↓      ↓
//    dfs(3)   dfs(4)  dfs(4)  dfs(5)
//       ↓        ↓      ↓      ↓
//    dfs(4)   dfs(5)  dfs(5)  dfs(6)

func rob(nums []int) int {
	n := len(nums)
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = max(dfs(i+1), nums[i]+dfs(i+2))
		return memo[i]
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
