package main

import "fmt"

func main() {
	var N int
	var S string

	fmt.Scan(&N, &S)

	first := S[0 : len(S)/2]
	last := S[len(S)/2:]

	if first == last {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

// Nを利用して早期リターンをすることができる
// package main

// import "fmt"

// func main() {
// 	var N int
// 	var S string

// 	fmt.Scan(&N, &S)

// 	if N%2 != 0 || N != len(S) {
// 		fmt.Println("No")
// 		return
// 	}

// 	first := S[:N/2]
// 	last := S[N/2:]
// 	if first == last {
// 		fmt.Println("Yes")
// 	} else {
// 		fmt.Println("No")
// 	}
// }

// ループを使った解き方

// package main

// import "fmt"

// func main() {
// 	var N int
// 	var S string
// 	fmt.Scan(&N, &S)

// 	// 偶数長でないなら不可能
// 	if N%2 != 0 || N != len(S) {
// 		fmt.Println("No")
// 		return
// 	}

// 	half := N / 2
// 	for i := 0; i < half; i++ {
// 		if S[i] != S[i+half] {
// 			fmt.Println("No")
// 			return
// 		}
// 	}
// 	fmt.Println("Yes")
// }
