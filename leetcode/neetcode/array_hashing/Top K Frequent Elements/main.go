package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, len(nums))

	for i := range nums {
		count[nums[i]]++
	}

	countArr := make([][2]int, 0, len(count))

	for num, cnt := range count {
		countArr = append(countArr, [2]int{num, cnt})
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i][1] > countArr[j][1]
	})

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = countArr[i][0]
	}

	return res
}

func main() {

	fmt.Println(topKFrequent([]int{1, 2, 2, 3, 3, 3}, 2))
	fmt.Println(topKFrequent([]int{7, 7}, 1))

}
