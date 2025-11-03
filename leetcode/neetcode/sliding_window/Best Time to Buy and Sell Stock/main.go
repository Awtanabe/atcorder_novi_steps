package main

import "fmt"

func maxProfit(prices []int) int {

	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}

	if max < 0 {
		return 0
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{110, 8, 7, 5, 2}))
}
