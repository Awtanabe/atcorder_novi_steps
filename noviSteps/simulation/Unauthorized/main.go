// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var N int
// 	fmt.Fscan(in, &N)

// 	loggedIn := false
// 	errors := 0

// 	for i := 0; i < N; i++ {
// 		var s string
// 		fmt.Fscan(in, &s)

// 		switch s {
// 		case "login":
// 			loggedIn = true
// 		case "logout":
// 			loggedIn = false
// 		case "public":
// 			// 何もしない
// 		case "private":
// 			if !loggedIn {
// 				errors++
// 			}
// 		}
// 	}
// 	fmt.Println(errors)
// }

package main

import "fmt"

func main() {
	var N int
	status := false

	fmt.Scan(&N)
	actives := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&actives[i])
	}

	c := 0

	for i := range actives {
		switch actives[i] {
		case "login":
			status = true
		case "logout":
			status = false
		case "private":
			if !status {
				c++
			}
		}
	}
	fmt.Print(c)
}
