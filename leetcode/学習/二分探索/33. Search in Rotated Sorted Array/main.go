package main

import "fmt"

func search(nums []int, target int) int {
	// left は 0 から始まる
	// right は末尾
	left, right := 0, len(nums)-1

	for left <= right {
		//真ん中
		mid := (left + right) / 2

		// ターゲットが真ん中ならretun
		if nums[mid] == target {
			return mid
		}

		// 左半分がソート済みの場合
		if nums[left] <= nums[mid] {
			// 左がわにあるなら
			if nums[left] <= target && target < nums[mid] {
				// 右端を 中央に移動する
				right = mid - 1
			} else {
				// 左側にないのなら lect
				left = mid + 1
			}
		} else {
			// 右半分がソート済みの場合
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
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1
}
