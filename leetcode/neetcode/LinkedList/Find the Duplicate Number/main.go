package main

import "fmt"

func findDuplicate(nums []int) int {

	tmp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmp[nums[i]]++
	}

	for k, v := range tmp {
		if v > 1 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 2, 2}))
}
