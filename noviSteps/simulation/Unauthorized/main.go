package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	loggedIn := false
	errors := 0

	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(in, &s)

		switch s {
		case "login":
			loggedIn = true
		case "logout":
			loggedIn = false
		case "public":
			// 何もしない
		case "private":
			if !loggedIn {
				errors++
			}
		}
	}
	fmt.Println(errors)
}
