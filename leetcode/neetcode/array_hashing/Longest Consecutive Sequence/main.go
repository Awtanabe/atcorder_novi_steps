package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. すべての数をセットに登録
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0

	// 2. 各数をチェック
	for _, num := range nums {
		// num-1 が存在するならスキップ（先頭じゃない）
		if _, exists := set[num-1]; exists {
			continue
		}

		// 3. ここから連続列をカウント
		length := 1
		for {
			_, exists := set[num+length]
			if !exists {
				break
			}
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

func main() {

	fmt.Println(longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5}))
}
