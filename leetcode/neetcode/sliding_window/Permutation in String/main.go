package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count1 := make([]int, 26)
	count2 := make([]int, 26)

	// s1 の文字頻度をカウント
	for _, c := range s1 {
		count1[c-'a']++
	}

	for i := 0; i < len(s2); i++ {
		count2[s2[i]-'a']++

		// ウィンドウサイズが s1 の長さを超えたら左端を削る
		if i >= len(s1) {
			count2[s2[i-len(s1)]-'a']--
		}

		// 二つの頻度配列を比較
		if equal(count1, count2) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkInclusion("abc", "lecabee"))  // true ("cab"が含まれる)
	fmt.Println(checkInclusion("abc", "lecaabee")) // false
}
