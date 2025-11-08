package main

import (
	"fmt"
)

// バックトラッキングは returnをすれば処理をそこまでで終了させることができる

// func combinationSum(nums []int, target int) [][]int {
// 	res := [][]int{}

// 	var dfs func(int, []int, int)
// 	dfs = func(i int, cur []int, total int) {
// 		if total == target {
// 			temp := make([]int, len(cur))
// 			copy(temp, cur)
// 			res = append(res, temp)
// 			return
// 		}
// 		if i >= len(nums) || total > target {
// 			return
// 		}

// 		// cur は[]intで選ぶか選ばないかをやってる
// 		cur = append(cur, nums[i])
// 		dfs(i, cur, total+nums[i])
// 		cur = cur[:len(cur)-1]
// 		dfs(i+1, cur, total)
// 	}

// 	dfs(0, []int{}, 0)
// 	return res
// }

func combinationSum(nums []int, target int) [][]int {

	res := [][]int{}

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {

		if target == total {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		if i >= len(nums) || target < total {
			return
		}

		cur = append(cur, nums[i])
		dfs(i, cur, total+nums[i])
		cur = cur[:len(cur)-1]
		dfs(i+1, cur, total)

	}
	dfs(0, []int{}, 0)

	return res
}

func main() {

	fmt.Println(combinationSum([]int{2, 3}, 7))
}
