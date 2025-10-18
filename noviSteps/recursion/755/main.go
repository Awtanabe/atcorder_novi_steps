package main

import (
	"fmt"
	"strconv"
)

// var N int
// var ans int

// func main() {
// 	fmt.Scan(&N)
// 	dfs(0)
// 	fmt.Println(ans)
// }

// func dfs(cur int) {
// 	if cur > N {
// 		return
// 	}
// 	if is753(cur) {
// 		ans++
// 	}
// 	for _, d := range []int{3, 5, 7} {
// 		dfs(cur*10 + d)
// 	}
// }

// func is753(x int) bool {
// 	if x == 0 {
// 		return false
// 	}
// 	s := strconv.Itoa(x)
// 	has3, has5, has7 := false, false, false
// 	for _, c := range s {
// 		if c == '3' {
// 			has3 = true
// 		} else if c == '5' {
// 			has5 = true
// 		} else if c == '7' {
// 			has7 = true
// 		}
// 	}
// 	return has3 && has5 && has7
// }

var N, ans int

func main() {
	fmt.Scan(&N)
	defs(0)
	fmt.Print(ans)
}

func is753(x int) bool {

	has3, has7, has5 := false, false, false

	ss := strconv.Itoa(x)

	for _, a := range ss {
		if a == '3' {
			has3 = true
		}
		if a == '5' {
			has5 = true
		}
		if a == '7' {
			has7 = true
		}
	}
	return has3 && has5 && has7
}

// 最近処理なので 全探索になる
func defs(cur int) {
	if cur > N {
		return
	}

	if is753(cur) {
		ans++
	}

	for _, d := range []int{3, 5, 7} {
		// N が400なら
		// 3 ~ 333 までを再帰処理で行っている
		defs(cur*10 + d)
	}
}

// 総和の再帰処理
// func recursion(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}

// 	return n + recursion(n-1)
// }

// func main() {
// 	re := recursion(10)
// 	fmt.Print(re)
// }
