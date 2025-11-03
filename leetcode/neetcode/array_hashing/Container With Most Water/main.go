package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		// 幅
		width := r - l
		// 高さ（低い方）
		h := height[l]
		if height[r] < h {
			h = height[r]
		}

		// 面積を計算
		area := width * h
		if area > maxArea {
			maxArea = area
		}

		// 高さの低い方を動かす
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6})) // 36
	fmt.Println(maxArea([]int{2, 2, 2}))                // 4
}
