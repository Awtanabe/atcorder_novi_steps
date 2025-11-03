package main

import "fmt"

func subsets(nums []int) [][]int {

	res := make([][]int, 0)

	subset := make([]int, 0)
	var dfs func(n int)
	dfs = func(n int) {

		if len(nums) == n {
			d := append([]int(nil), subset...)
			res = append(res, d)
			return
		}

		// 入れる ⭐️ここ惜しい
		subset = append(subset, nums[n])
		dfs(n + 1)

		// 差し引く
		subset = subset[:len(subset)-1]
		dfs(n + 1)
	}

	dfs(0)

	return res
}

func main() {

	fmt.Println(subsets([]int{1, 2, 3}))
}
