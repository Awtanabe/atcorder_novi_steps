package main

import "fmt"

func main() {
	var H, W, d int

	fmt.Scan(&H, &W)

	data := make([][]int, H)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Scan(&d)
			data[i] = append(data[i], d)
		}
	}

	fmt.Print(data)
}
