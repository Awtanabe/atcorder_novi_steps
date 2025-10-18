package main

import "fmt"

// index か -1
func binarySeach(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {

	fmt.Println(binarySeach([]int{0, 1, 2, 4, 5, 6, 7}, 3))

}
