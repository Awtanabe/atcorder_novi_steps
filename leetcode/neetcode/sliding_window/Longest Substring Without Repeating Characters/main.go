// func lengthOfLongestSubstring(s string) int {
// 	max := 1
// 	for i := 0; i < len(s); i++ {
// 		tmp := 1
// 		for j := i + 1; j < len(s); j++ {
// 			if s[j]-s[i] == 1 && right {
// 				tmp++
// 				if max < tmp {
// 					max = tmp
// 				}
// 			}
// 		}
// 	}

// 	return max
// }

// func main() {
// 	// abcabcbb
// 	// pwwkew
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// }

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	seen := make(map[byte]int)

	maxlen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := seen[s[right]]; ok && idx >= left {
			left = idx + 1
		}

		seen[s[right]] = right

		// いまいちここらへんは分からない
		if (right-left)+1 > maxlen {
			maxlen = right - left + 1
		}
	}
	return maxlen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // → 3 ("wke")
}
