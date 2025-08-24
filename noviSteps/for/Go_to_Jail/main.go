package main

import "fmt"

// スライスのインデックスを使った利用
func main() {
	var N, a, b, c int

	fmt.Scan(&N)

	data := make([][]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a, &b)
		data[i] = []int{a, b}
	}

	for _, row := range data {
		if row[0] == row[1] {
			c++
		}

		if c >= 3 {
			break
		}

		if row[0] != row[1] {
			c = 0
		}
	}

	if c >= 3 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

// package main

// import "fmt"

// func main() {
// 	var N, a, b, count int

// 	fmt.Scan(&N)

// 	for i := 0; i < N; i++ {
// 		fmt.Scan(&a, &b)
// 		if a == b {
// 			count++
// 		} else if count < 3 {
// 			count = 0
// 		}
// 	}

// 	if count >= 3 {
// 		fmt.Print("Yes")
// 	} else {
// 		fmt.Print("No")
// 	}
// }
