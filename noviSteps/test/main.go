// package main

// import "fmt"

// func main() {
// 	var N, X int
// 	var Y string

// 	fmt.Scan(&N)

// 	data := make([]string, N)

// 	for i := 0; i < N; i++ {
// 		fmt.Scan(&data[i])
// 	}

// 	fmt.Scan(&X, &Y)

// 	if data[X-1] == Y {
// 		fmt.Print("Yes")
// 	} else {
// 		fmt.Print("No")
// 	}
// }

package main

import "fmt"

func main() {
	var N, c int
	var S string

	fmt.Scan(&N)
	fmt.Scan(&S)

	for i := 0; i < (2*N)-1; i++ {
		s := []rune(S)
		if s[i] == s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			c++
		}
	}

	fmt.Print(c)
}
