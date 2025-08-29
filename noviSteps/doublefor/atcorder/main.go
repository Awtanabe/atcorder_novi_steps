package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)

	isACGT := func(b byte) bool {
		return b == 'A' || b == 'C' || b == 'G' || b == 'T'
	}

	best, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if isACGT(s[i]) {
			cur++
			if cur > best {
				best = cur
			}
		} else {
			cur = 0
		}
	}
	fmt.Println(best)
}
