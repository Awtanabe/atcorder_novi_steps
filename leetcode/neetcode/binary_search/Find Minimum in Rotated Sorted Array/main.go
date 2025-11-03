package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {

	// fmt.Println(findMin([]int{4, 5, 6, 7}))

	fmt.Println(findMin([]int{3, 4, 5, 6, 1, 2}))
}
