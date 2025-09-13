// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	var s string
// 	fmt.Fscan(in, &s)

// 	isACGT := func(b byte) bool {
// 		return b == 'A' || b == 'C' || b == 'G' || b == 'T'
// 	}

// 	best, cur := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		if isACGT(s[i]) {
// 			cur++
// 			if cur > best {
// 				best = cur
// 			}
// 		} else {
// 			cur = 0
// 		}
// 	}
// 	fmt.Println(best)
// }

package main

import "fmt"

func main() {
	var S string

	fmt.Scan(&S)
	max := 0

	tmp := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'A' || S[i] == 'C' || S[i] == 'G' || S[i] == 'T' {
			tmp += 1
		} else {
			if max < tmp {
				max = tmp
			}
			tmp = 0
		}
	}

	// ここの一度もelse に入らないケースの考慮が抜けている
	// ⭐️テストが甘い 閾値、最大、最小の3つはテストをする癖を
	if max < tmp {
		max = tmp
	}

	fmt.Print(max)
}
