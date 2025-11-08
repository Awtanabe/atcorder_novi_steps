package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	// sort必要
	sort.Ints(candidates)

	var dfs func(int, []int, int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if total > target || i == len(candidates) {
			return
		}

		cur = append(cur, candidates[i])
		// ここ + 1
		dfs(i+1, cur, total+candidates[i])
		cur = cur[:len(cur)-1]

		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, cur, total)
	}

	dfs(0, []int{}, 0)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{9, 2, 2, 4, 6, 1, 5}, 8))
}
