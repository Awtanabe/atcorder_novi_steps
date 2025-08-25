package main

import "fmt"

func main() {
	var H, W, d int

	fmt.Scan(&H, &W)

	data := make([][]int, H)

	for i := 0; i < H; i++ {
		data[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Scan(&d)
			data[i][j] = d
		}
	}

	for _, row := range data {
		for _, d := range row {
			if d == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(rune('A' + d - 1)))

			}
		}
		fmt.Println("")
	}
}
