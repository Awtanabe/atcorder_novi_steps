package main

import "fmt"

// func rob(nums []int) int {

// 	if 2 == len(nums) {

// 		if nums[0] < nums[1] {
// 			return nums[1]
// 		} else {
// 			return nums[0]
// 		}
// 	}

// 	max := 0
// 	tmp := 0
// 	var rr func(i int)
// 	rr = func(i int) {
// 		if i > len(nums)-1 {
// 			return
// 		}

// 		tmp += nums[i]
// 		if tmp > max {
// 			max = tmp
// 		}

// 		rr(i + 2)
// 	}
// 	rr(0)
// 	return max
// }

// 再帰処理
func rob(nums []int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		return max(dfs(i+1), nums[i]+dfs(i+2))
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
	fmt.Println(rob([]int{1, 2}))
	// fmt.Println(rob([]int{1, 1, 3, 3}))
	// fmt.Println(rob([]int{2, 9, 8, 3, 6}))
}
