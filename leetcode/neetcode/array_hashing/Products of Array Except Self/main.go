package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {

		tmp := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				tmp *= nums[j]
			}
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 4, 6}))
}
