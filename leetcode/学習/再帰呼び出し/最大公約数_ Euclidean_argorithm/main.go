package main

import "fmt"

func aaa(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}

	return aaa(b, r)
}

func main() {

	fmt.Println(aaa(18, 12))
}
