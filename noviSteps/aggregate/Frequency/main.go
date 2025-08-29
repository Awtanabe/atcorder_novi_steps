package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	var cnt [26]int

	for i := range S {
		cnt[S[i]-'a']++
	}

	max := 0

	for i := range cnt {
		if cnt[i] > cnt[max] {
			max = i
		}
	}

	fmt.Print(max)
}

// mapパターン
// package main

// import "fmt"

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	count := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		count[s[i]]++
// 	}

// 	maxCnt := -1
// 	var ans byte = 'a'
// 	for ch := byte('a'); ch <= 'z'; ch++ {
// 		if count[ch] > maxCnt {
// 			maxCnt = count[ch]
// 			ans = ch
// 		}
// 	}
// 	fmt.Printf("%c\n", ans)
// }

// バケットパターン
// package main

// import "fmt"

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	var cnt [26]int
// 	for i := 0; i < len(s); i++ {
// 		cnt[s[i]-'a']++
// 	}

// 	best := 0
// 	for i := 1; i < 26; i++ {
// 		if cnt[i] > cnt[best] {
// 			best = i
// 		}
// 	}
// 	fmt.Printf("%c\n", 'a'+byte(best))
// }
