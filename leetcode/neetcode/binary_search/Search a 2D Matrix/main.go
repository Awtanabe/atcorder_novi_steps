package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(row []int, target int) bool {

	left, right := 0, len(row)-1

	for left <= right {
		mid := (left + right) / 2

		if row[mid] == target {
			return true
		}

		if row[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 2, 4, 8},
		{0, 11, 12, 13},
		{14, 20, 30, 40},
	}, 15))
}
