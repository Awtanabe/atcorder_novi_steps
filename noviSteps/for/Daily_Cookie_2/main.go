package main

import "fmt"

func main() {
	var N, D, i int
	var S string

	fmt.Scan(&N, &D)
	fmt.Scan(&S)

	s := []rune(S)

	for 0 != D {
		if s[N-1-i] == '@' {
			s[N-1-i] = '.'
			D--
		}
		i++
	}

	fmt.Println(string(s))
}
