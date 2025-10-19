package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// 右側が小さい → 最小値は右側にある
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// 右側のほうが大きい → 最小値は左を含む
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
}
