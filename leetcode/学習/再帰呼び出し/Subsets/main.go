// 問題
// https://interviewcat.dev/p/coding-interviewcat/recursion-backtracking-backtracking
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 部分配列を作れ

package main

import "fmt"

// ⭐️subsetのスライスは、スタックごとに保存される
// 処理の流れ rec(1) => rec(3)までいったら 選ばない処理に入る
// return 今呼び出されている関数は終了する
// rec(3) が終わったら re(2)が再開する

// 流れ

// rec(0): subset=[]
//   ├─ 1を選ぶ → rec(1): subset=[1]
//   │    ├─ 2を選ぶ → rec(2): subset=[1,2]
//   │    │    ├─ 3を選ぶ → rec(3): subset=[1,2,3]  ✅追加
//   │    │    └─ 3を選ばない → rec(3): subset=[1,2] ✅追加
//   │    └─ 2を選ばない → rec(2): subset=[1]
//   │         ├─ 3を選ぶ → rec(3): subset=[1,3] ✅追加
//   │         └─ 3を選ばない → rec(3): subset=[1] ✅追加
//   └─ 1を選ばない → rec(1): subset=[]
//        ├─ 2を選ぶ → rec(2): subset=[2]
//        │    ├─ 3を選ぶ → rec(3): subset=[2,3] ✅追加
//        │    └─ 3を選ばない → rec(3): subset=[2] ✅追加
//        └─ 2を選ばない → rec(2): subset=[]
//             ├─ 3を選ぶ → rec(3): subset=[3] ✅追加
//             └─ 3を選ばない → rec(3): subset=[] ✅追加

func subsets(nums []int) [][]int {
	var ans [][]int
	var subset []int
	n := len(nums)

	var rec func(i int)
	rec = func(i int) {
		// 現在の状態を出力
		fmt.Printf("rec(%d) 呼び出し中: subset=%v i=%d n=%d\n", i, subset, i, n)

		// 配列を全部見終わったら現在のsubsetを答えに追加
		if i == n {
			cp := append([]int(nil), subset...)
			ans = append(ans, cp)
			fmt.Printf("  → 末尾に到達: subset=%v を追加\n", cp)
			return // 今呼び出されている関数は終了する
		}

		// (1) nums[i]を選ぶ
		fmt.Printf("  nums[%d]=%d を選ぶ\n", i, nums[i])
		subset = append(subset, nums[i])
		rec(i + 1)

		// (2) nums[i]を選ばない
		fmt.Printf("  nums[%d]=%d を選ばない (popして戻す)\n", i, nums[i])
		// 元に戻す動作
		subset = subset[:len(subset)-1]
		rec(i + 1)
	}

	rec(0)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println("\n最終結果:", result)
}
