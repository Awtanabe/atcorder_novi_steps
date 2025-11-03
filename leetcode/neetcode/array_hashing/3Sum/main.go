package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	tmp := map[[3]int]struct{}{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	var res [][]int
	for key, _ := range tmp {
		res = append(res, []int{key[0], key[1], key[2]})
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
