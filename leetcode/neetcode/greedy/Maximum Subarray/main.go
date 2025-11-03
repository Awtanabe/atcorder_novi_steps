package main

import (
	"fmt"
)

// func maxSubArray(nums []int) int {

// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	max := 0

// 	for i := 0; i < len(nums); i++ {
// 		tmp := 0
// 		tmp += nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			tmp += nums[j]
// 			if max <= tmp {
// 				max = tmp
// 			}

// 		}
// 	}

// 	return max
// }

func maxSubArray(nums []int) int {
	res := nums[0]

	var dfs func(i, cur int)
	dfs = func(i, cur int) {
		if i == len(nums) {
			return
		}

		cur = max(nums[i], cur+nums[i]) // 連続するか、ここでリセットするか
		res = max(res, cur)
		dfs(i+1, cur)
	}

	dfs(0, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(maxSubArray([]int{2, -3, 4, -2, 2, 1, -1, 4}))
	fmt.Println(maxSubArray([]int{-2, 1}))
}
