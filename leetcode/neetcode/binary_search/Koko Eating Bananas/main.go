package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	speed := 1
	for {
		totalTime := 0
		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(speed)))
		}

		if totalTime <= h {
			return speed
		}
		speed += 1
	}
	return speed
}
func main() {
	// fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	// fmt.Println(minEatingSpeed([]int{1, 4, 3, 2}, 9))
	fmt.Println(minEatingSpeed([]int{25, 10, 23, 4}, 4))
}
