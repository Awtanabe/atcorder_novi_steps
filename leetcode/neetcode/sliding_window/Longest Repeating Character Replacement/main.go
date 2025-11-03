package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left := 0
	maxCount := 0 // 現在のウィンドウ内で最も多い文字の出現回数
	res := 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		fmt.Println("right", count, right, string(s[right]), count[s[right]])
		if count[s[right]] > maxCount {
			maxCount = count[s[right]]
		}

		// 置換必要数がkを超えたら、ウィンドウを縮める
		// 必要な置換数 = (ウィンドウの長さ) - (最頻文字の数)
		// (right - left + 1) - maxCount
		for (right-left+1)-maxCount > k {
			count[s[left]]--
			left++
		}

		// 最大長を更新
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

func main() {
	fmt.Println(characterReplacement("XYYX", 2))    // 4
	fmt.Println(characterReplacement("AAABABB", 1)) // 5
}

// neet code

// func characterReplacement(s string, k int) int {
//     res := 0
//     charSet := make(map[byte]bool)

//     for i := 0; i < len(s); i++ {
//         charSet[s[i]] = true
//     }

//     for c := range charSet {
//         count, l := 0, 0
//         for r := 0; r < len(s); r++ {
//             if s[r] == c {
//                 count++
//             }

//             for (r - l + 1) - count > k {
//                 if s[l] == c {
//                     count--
//                 }
//                 l++
//             }

//             res = max(res, r - l + 1)
//         }
//     }

//     return res
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
