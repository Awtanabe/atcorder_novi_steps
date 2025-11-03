package main

import "fmt"

// func search(nums []int, target int) int {
// 	left, right := 0, len(nums)-1

// 	for left < right {

// 		mid := (left + right) / 2
// 		if nums[mid] == target {
// 			return mid
// 		}

// 		if nums[mid] > target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return -1
// }

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 左側がソート済み
		// target 比較してない
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右側がソート済み
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	fmt.Println(search([]int{3, 4, 5, 6, 1, 2}, 1))
}
