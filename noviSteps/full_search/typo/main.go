package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	for i := 0; i < len(S)-1; i++ {
		if S == T {
			fmt.Print("Yes")
			return
		}
		s := []rune(S)
		s[i], s[i+1] = s[i+1], s[i]
		if string(s) == T {
			fmt.Print("Yes")
			return
		}
	}
	fmt.Print("No")
}
