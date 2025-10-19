package main

import "fmt"

func hasDuplicate(nums []int) bool {

	memo := make(map[int]bool)

	for i := range nums {
		if _, ok := memo[nums[i]]; ok {
			return true
		} else {
			memo[nums[i]] = true
		}
	}
	return false

}

func main() {
	fmt.Println(hasDuplicate([]int{1, 2, 3, 4}))
}
