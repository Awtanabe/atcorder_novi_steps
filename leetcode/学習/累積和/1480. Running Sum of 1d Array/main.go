package main

import "fmt"

func runningSum(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = nums[i]
		} else {
			res[i] = nums[i] + res[i-1]
		}
	}
	return res
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))

}
