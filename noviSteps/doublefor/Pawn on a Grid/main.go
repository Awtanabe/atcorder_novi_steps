package main

import "fmt"

func main() {
	var H, W, c int

	fmt.Scan(&H, &W)
	data := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scan(&data[i])
	}

	for _, d := range data {

		for _, cc := range d {
			if cc == '#' {
				c++
			}
		}
	}

	fmt.Print(c)
}
